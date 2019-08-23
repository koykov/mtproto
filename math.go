package mtproto

import (
	"crypto/aes"
	sha1lib "crypto/sha1"
	"errors"
	"math/big"
	"math/rand"
	"time"
)

const (
	telegramPublicKeyFP = 14101943622620965665
)

func sha1(data []byte) []byte {
	r := sha1lib.Sum(data)
	return r[:]
}

func doRSAencrypt(em []byte) []byte {
	z := make([]byte, 255)
	copy(z, em)

	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(z), big.NewInt(int64(conf.RsaKey.E)), conf.RsaKey.N)

	res := make([]byte, 256)
	copy(res, c.Bytes())

	return res
}

func splitPQ(pq *big.Int) (p1, p2 *big.Int) {
	value0 := big.NewInt(0)
	value1 := big.NewInt(1)
	value15 := big.NewInt(15)
	value17 := big.NewInt(17)
	rndmax := big.NewInt(0).SetBit(big.NewInt(0), 64, 1)

	what := big.NewInt(0).Set(pq)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := big.NewInt(0)
	i := 0
	for !(g.Cmp(value1) == 1 && g.Cmp(what) == -1) {
		q := big.NewInt(0).Rand(rnd, rndmax)
		q = q.And(q, value15)
		q = q.Add(q, value17)
		q = q.Mod(q, what)

		x := big.NewInt(0).Rand(rnd, rndmax)
		whatnext := big.NewInt(0).Sub(what, value1)
		x = x.Mod(x, whatnext)
		x = x.Add(x, value1)

		y := big.NewInt(0).Set(x)
		lim := 1 << (uint(i) + 18)
		j := 1
		flag := true

		for j < lim && flag {
			a := big.NewInt(0).Set(x)
			b := big.NewInt(0).Set(x)
			c := big.NewInt(0).Set(q)

			for b.Cmp(value0) == 1 {
				b2 := big.NewInt(0)
				if b2.And(b, value1).Cmp(value0) == 1 {
					c.Add(c, a)
					if c.Cmp(what) >= 0 {
						c.Sub(c, what)
					}
				}
				a.Add(a, a)
				if a.Cmp(what) >= 0 {
					a.Sub(a, what)
				}
				b.Rsh(b, 1)
			}
			x.Set(c)

			z := big.NewInt(0)
			if x.Cmp(y) == -1 {
				z.Add(what, x)
				z.Sub(z, y)
			} else {
				z.Sub(x, y)
			}
			g.GCD(nil, nil, z, what)

			if (j & (j - 1)) == 0 {
				y.Set(x)
			}
			j = j + 1

			if g.Cmp(value1) != 0 {
				flag = false
			}
		}
		i = i + 1
	}

	p1 = big.NewInt(0).Set(g)
	p2 = big.NewInt(0).Div(what, g)

	if p1.Cmp(p2) == 1 {
		p1, p2 = p2, p1
	}

	return
}

func makeGAB(g int32, ga, dhPrime *big.Int) (b, gb, gab *big.Int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndmax := big.NewInt(0).SetBit(big.NewInt(0), 2048, 1)
	b = big.NewInt(0).Rand(rnd, rndmax)
	gb = big.NewInt(0).Exp(big.NewInt(int64(g)), b, dhPrime)
	gab = big.NewInt(0).Exp(ga, b, dhPrime)

	return
}

func generateAES(msgKey, authKey []byte, decode bool) ([]byte, []byte) {
	var x int
	if decode {
		x = 8
	} else {
		x = 0
	}
	aesKey := make([]byte, 0, 32)
	aesIv := make([]byte, 0, 32)
	ta := make([]byte, 0, 48)
	tb := make([]byte, 0, 48)
	tc := make([]byte, 0, 48)
	td := make([]byte, 0, 48)

	ta = append(ta, msgKey...)
	ta = append(ta, authKey[x:x+32]...)

	tb = append(tb, authKey[32+x:32+x+16]...)
	tb = append(tb, msgKey...)
	tb = append(tb, authKey[48+x:48+x+16]...)

	tc = append(tc, authKey[64+x:64+x+32]...)
	tc = append(tc, msgKey...)

	td = append(td, msgKey...)
	td = append(td, authKey[96+x:96+x+32]...)

	sha1A := sha1(ta)
	sha1B := sha1(tb)
	sha1C := sha1(tc)
	sha1D := sha1(td)

	aesKey = append(aesKey, sha1A[0:8]...)
	aesKey = append(aesKey, sha1B[8:8+12]...)
	aesKey = append(aesKey, sha1C[4:4+12]...)

	aesIv = append(aesIv, sha1A[8:8+12]...)
	aesIv = append(aesIv, sha1B[0:8]...)
	aesIv = append(aesIv, sha1C[16:16+4]...)
	aesIv = append(aesIv, sha1D[0:8]...)

	return aesKey, aesIv
}

func doAES256IGEEncrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("AES256IGE: data too small to encrypt")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("AES256IGE: data not divisible by block size")
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, iv[:aes.BlockSize])
	copy(y, iv[aes.BlockSize:])
	encrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(x, data[i:i+aes.BlockSize])
		block.Encrypt(t, x)
		xor(t, y)
		x, y = t, data[i:i+aes.BlockSize]
		copy(encrypted[i:], t)
		i += aes.BlockSize
	}

	return encrypted, nil
}

func doAES256IGEDecrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("AES256IGE: data too small to decrypt")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("AES256IGE: data not divisible by block size")
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, iv[:aes.BlockSize])
	copy(y, iv[aes.BlockSize:])
	decrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(y, data[i:i+aes.BlockSize])
		block.Decrypt(t, y)
		xor(t, x)
		y, x = t, data[i:i+aes.BlockSize]
		copy(decrypted[i:], t)
		i += aes.BlockSize
	}

	return decrypted, nil

}

func xor(dst, src []byte) {
	for i := range dst {
		dst[i] = dst[i] ^ src[i]
	}
}
