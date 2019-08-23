package mtproto

import "fmt"

const (
	CrcBoolFalse                              = 0xbc799737
	CrcBoolTrue                               = 0x997275b5
	CrcVector                                 = 0x1cb5c415
	CrcError                                  = 0xc4b9f9bb
	CrcNull                                   = 0x56730bcc
	CrcInputPeerEmpty                         = 0x7f3b18ea
	CrcInputPeerSelf                          = 0x7da07ec9
	CrcInputPeerContact                       = 0x1023dbe8
	CrcInputPeerForeign                       = 0x9b447325
	CrcInputPeerChat                          = 0x179be863
	CrcInputUserEmpty                         = 0xb98886cf
	CrcInputUserSelf                          = 0xf7c1b13f
	CrcInputUserContact                       = 0x86e94f65
	CrcInputUserForeign                       = 0x655e74ff
	CrcInputPhoneContact                      = 0xf392b7f4
	CrcInputFile                              = 0xf52ff27f
	CrcInputMediaEmpty                        = 0x9664f57f
	CrcInputMediaUploadedPhoto                = 0x2dc53a7d
	CrcInputMediaPhoto                        = 0x8f2ab2ec
	CrcInputMediaGeoPoint                     = 0xf9c44144
	CrcInputMediaContact                      = 0xa6e45987
	CrcInputMediaUploadedVideo                = 0x133ad6f6
	CrcInputMediaUploadedThumbVideo           = 0x9912dabf
	CrcInputMediaVideo                        = 0x7f023ae6
	CrcInputChatPhotoEmpty                    = 0x1ca48f57
	CrcInputChatUploadedPhoto                 = 0x94254732
	CrcInputChatPhoto                         = 0xb2e1bf08
	CrcInputGeoPointEmpty                     = 0xe4c123d6
	CrcInputGeoPoint                          = 0xf3b7acc9
	CrcInputPhotoEmpty                        = 0x1cd7bf0d
	CrcInputPhoto                             = 0xfb95c6c4
	CrcInputVideoEmpty                        = 0x5508ec75
	CrcInputVideo                             = 0xee579652
	CrcInputFileLocation                      = 0x14637196
	CrcInputVideoFileLocation                 = 0x3d0364ec
	CrcInputPhotoCropAuto                     = 0xade6b004
	CrcInputPhotoCrop                         = 0xd9915325
	CrcInputAppEvent                          = 0x770656a8
	CrcPeerUser                               = 0x9db1bc6d
	CrcPeerChat                               = 0xbad0e5bb
	CrcStorageFileUnknown                     = 0xaa963b05
	CrcStorageFileJpeg                        = 0x007efe0e
	CrcStorageFileGif                         = 0xcae1aadf
	CrcStorageFilePng                         = 0x0a4f63c0
	CrcStorageFilePdf                         = 0xae1e508d
	CrcStorageFileMp3                         = 0x528a0677
	CrcStorageFileMov                         = 0x4b09ebbc
	CrcStorageFilePartial                     = 0x40bc6f52
	CrcStorageFileMp4                         = 0xb3cea0e4
	CrcStorageFileWebp                        = 0x1081464c
	CrcFileLocationUnavailable                = 0x7c596b46
	CrcFileLocation                           = 0x53d69076
	CrcUserEmpty                              = 0x200250ba
	CrcUserSelf                               = 0x7007b451
	CrcUserContact                            = 0xcab35e18
	CrcUserRequest                            = 0xd9ccc4ef
	CrcUserForeign                            = 0x075cf7a8
	CrcUserDeleted                            = 0xd6016d7a
	CrcUserProfilePhotoEmpty                  = 0x4f11bae1
	CrcUserProfilePhoto                       = 0xd559d8c8
	CrcUserStatusEmpty                        = 0x09d05049
	CrcUserStatusOnline                       = 0xedb93949
	CrcUserStatusOffline                      = 0x008c703f
	CrcChatEmpty                              = 0x9ba2d800
	CrcChat                                   = 0x6e9c9bc7
	CrcChatForbidden                          = 0xfb0ccc41
	CrcChatFull                               = 0x630e61be
	CrcChatParticipant                        = 0xc8d7493e
	CrcChatParticipantsForbidden              = 0x0fd2bb8a
	CrcChatParticipants                       = 0x7841b415
	CrcChatPhotoEmpty                         = 0x37c1011c
	CrcChatPhoto                              = 0x6153276a
	CrcMessageEmpty                           = 0x83e5de54
	CrcMessage                                = 0x567699b3
	CrcMessageForwarded                       = 0xa367e716
	CrcMessageService                         = 0x1d86f70e
	CrcMessageMediaEmpty                      = 0x3ded6320
	CrcMessageMediaPhoto                      = 0xc8c45a2a
	CrcMessageMediaVideo                      = 0xa2d24290
	CrcMessageMediaGeo                        = 0x56e0d474
	CrcMessageMediaContact                    = 0x5e7d2f39
	CrcMessageMediaUnsupported                = 0x29632a36
	CrcMessageActionEmpty                     = 0xb6aef7b0
	CrcMessageActionChatCreate                = 0xa6638b9a
	CrcMessageActionChatEditTitle             = 0xb5a1ce5a
	CrcMessageActionChatEditPhoto             = 0x7fcb13a8
	CrcMessageActionChatDeletePhoto           = 0x95e3fbef
	CrcMessageActionChatAddUser               = 0x5e3cfc4b
	CrcMessageActionChatDeleteUser            = 0xb2ae9b0c
	CrcDialog                                 = 0xab3a99ac
	CrcPhotoEmpty                             = 0x2331b22d
	CrcPhoto                                  = 0x22b56751
	CrcPhotoSizeEmpty                         = 0x0e17e23c
	CrcPhotoSize                              = 0x77bfb61b
	CrcPhotoCachedSize                        = 0xe9a734fa
	CrcVideoEmpty                             = 0xc10658a8
	CrcVideo                                  = 0x388fa391
	CrcGeoPointEmpty                          = 0x1117dd5f
	CrcGeoPoint                               = 0x2049d70c
	CrcAuthCheckedPhone                       = 0xe300cc3b
	CrcAuthSentCode                           = 0xefed51d9
	CrcAuthAuthorization                      = 0xf6b673a4
	CrcAuthExportedAuthorization              = 0xdf969c2d
	CrcInputNotifyPeer                        = 0xb8bc5b0c
	CrcInputNotifyUsers                       = 0x193b4417
	CrcInputNotifyChats                       = 0x4a95e84e
	CrcInputNotifyAll                         = 0xa429b886
	CrcInputPeerNotifyEventsEmpty             = 0xf03064d8
	CrcInputPeerNotifyEventsAll               = 0xe86a2c74
	CrcInputPeerNotifySettings                = 0x46a2ce98
	CrcPeerNotifyEventsEmpty                  = 0xadd53cb3
	CrcPeerNotifyEventsAll                    = 0x6d1ded88
	CrcPeerNotifySettingsEmpty                = 0x70a68512
	CrcPeerNotifySettings                     = 0x8d5e11ee
	CrcWallPaper                              = 0xccb03657
	CrcUserFull                               = 0x771095da
	CrcContact                                = 0xf911c994
	CrcImportedContact                        = 0xd0028438
	CrcContactBlocked                         = 0x561bc879
	CrcContactSuggested                       = 0x3de191a1
	CrcContactStatus                          = 0xd3680c61
	CrcChatLocated                            = 0x3631cf4c
	CrcContactsForeignLinkUnknown             = 0x133421f8
	CrcContactsForeignLinkRequested           = 0xa7801f47
	CrcContactsForeignLinkMutual              = 0x1bea8ce1
	CrcContactsMyLinkEmpty                    = 0xd22a1c60
	CrcContactsMyLinkRequested                = 0x6c69efee
	CrcContactsMyLinkContact                  = 0xc240ebd9
	CrcContactsLink                           = 0xeccea3f5
	CrcContactsContactsNotModified            = 0xb74ba9d2
	CrcContactsContacts                       = 0x6f8b8cb2
	CrcContactsImportedContacts               = 0xad524315
	CrcContactsBlocked                        = 0x1c138d15
	CrcContactsBlockedSlice                   = 0x900802a1
	CrcContactsSuggested                      = 0x5649dcc5
	CrcMessagesDialogs                        = 0x15ba6c40
	CrcMessagesDialogsSlice                   = 0x71e094f3
	CrcMessagesMessages                       = 0x8c718e87
	CrcMessagesMessagesSlice                  = 0x0b446ae3
	CrcMessagesMessageEmpty                   = 0x3f4e0648
	CrcMessagesStatedMessages                 = 0x969478bb
	CrcMessagesStatedMessage                  = 0xd07ae726
	CrcMessagesSentMessage                    = 0xd1f4d35c
	CrcMessagesChats                          = 0x8150cbd8
	CrcMessagesChatFull                       = 0xe5d7d19c
	CrcMessagesAffectedHistory                = 0xb7de36f2
	CrcInputMessagesFilterEmpty               = 0x57e2f66c
	CrcInputMessagesFilterPhotos              = 0x9609a51c
	CrcInputMessagesFilterVideo               = 0x9fc00e65
	CrcInputMessagesFilterPhotoVideo          = 0x56e9f0e4
	CrcInputMessagesFilterPhotoVideoDocuments = 0xd95e73bb
	CrcInputMessagesFilterDocument            = 0x9eddf188
	CrcInputMessagesFilterAudio               = 0xcfc87522
	CrcUpdateNewMessage                       = 0x013abdb3
	CrcUpdateMessageID                        = 0x4e90bfd6
	CrcUpdateReadMessages                     = 0xc6649e31
	CrcUpdateDeleteMessages                   = 0xa92bfe26
	CrcUpdateUserTyping                       = 0x5c486927
	CrcUpdateChatUserTyping                   = 0x9a65ea1f
	CrcUpdateChatParticipants                 = 0x07761198
	CrcUpdateUserStatus                       = 0x1bfbd823
	CrcUpdateUserName                         = 0xa7332b73
	CrcUpdateUserPhoto                        = 0x95313b0c
	CrcUpdateContactRegistered                = 0x2575bbb9
	CrcUpdateContactLink                      = 0x51a48a9a
	CrcUpdateNewAuthorization                 = 0x8f06529a
	CrcUpdatesState                           = 0xa56c2a3e
	CrcUpdatesDifferenceEmpty                 = 0x5d75a138
	CrcUpdatesDifference                      = 0x00f49ca0
	CrcUpdatesDifferenceSlice                 = 0xa8fb1981
	CrcUpdatesTooLong                         = 0xe317af7e
	CrcUpdateShortMessage                     = 0xd3f45784
	CrcUpdateShortChatMessage                 = 0x2b2fbd4e
	CrcUpdateShort                            = 0x78d4dec1
	CrcUpdatesCombined                        = 0x725b04c3
	CrcUpdates                                = 0x74ae4240
	CrcPhotosPhotos                           = 0x8dca6aa5
	CrcPhotosPhotosSlice                      = 0x15051f54
	CrcPhotosPhoto                            = 0x20212ca8
	CrcUploadFile                             = 0x096a18d5
	CrcDcOption                               = 0x2ec2a43c
	CrcConfig                                 = 0x7dae33e0
	CrcNearestDc                              = 0x8e1a1775
	CrcHelpAppUpdate                          = 0x8987f311
	CrcHelpNoAppUpdate                        = 0xc45a6536
	CrcHelpInviteText                         = 0x18cb9f78
	CrcMessagesStatedMessagesLinks            = 0x3e74f5c6
	CrcMessagesStatedMessageLink              = 0xa9af2881
	CrcMessagesSentMessageLink                = 0xe9db4a3f
	CrcInputGeoChat                           = 0x74d456fa
	CrcInputNotifyGeoChatPeer                 = 0x4d8ddec8
	CrcGeoChat                                = 0x75eaea5a
	CrcGeoChatMessageEmpty                    = 0x60311a9b
	CrcGeoChatMessage                         = 0x4505f8e1
	CrcGeoChatMessageService                  = 0xd34fa24e
	CrcGeochatsStatedMessage                  = 0x17b1578b
	CrcGeochatsLocated                        = 0x48feb267
	CrcGeochatsMessages                       = 0xd1526db1
	CrcGeochatsMessagesSlice                  = 0xbc5863e8
	CrcMessageActionGeoChatCreate             = 0x6f038ebc
	CrcMessageActionGeoChatCheckin            = 0x0c7d53de
	CrcUpdateNewGeoChatMessage                = 0x5a68e3f7
	CrcWallPaperSolid                         = 0x63117f24
	CrcUpdateNewEncryptedMessage              = 0x12bcbd9a
	CrcUpdateEncryptedChatTyping              = 0x1710f156
	CrcUpdateEncryption                       = 0xb4a2e88d
	CrcUpdateEncryptedMessagesRead            = 0x38fe25b7
	CrcEncryptedChatEmpty                     = 0xab7ec0a0
	CrcEncryptedChatWaiting                   = 0x3bf703dc
	CrcEncryptedChatRequested                 = 0xc878527e
	CrcEncryptedChat                          = 0xfa56ce36
	CrcEncryptedChatDiscarded                 = 0x13d6dd27
	CrcInputEncryptedChat                     = 0xf141b5e1
	CrcEncryptedFileEmpty                     = 0xc21f497e
	CrcEncryptedFile                          = 0x4a70994c
	CrcInputEncryptedFileEmpty                = 0x1837c364
	CrcInputEncryptedFileUploaded             = 0x64bd0306
	CrcInputEncryptedFile                     = 0x5a17b5e5
	CrcInputEncryptedFileLocation             = 0xf5235d55
	CrcEncryptedMessage                       = 0xed18c118
	CrcEncryptedMessageService                = 0x23734b06
	CrcMessagesDhConfigNotModified            = 0xc0e24635
	CrcMessagesDhConfig                       = 0x2c221edd
	CrcMessagesSentEncryptedMessage           = 0x560f8935
	CrcMessagesSentEncryptedFile              = 0x9493ff32
	CrcInputFileBig                           = 0xfa4f0bb5
	CrcInputEncryptedFileBigUploaded          = 0x2dc173c8
	CrcUpdateChatParticipantAdd               = 0x3a0eeb22
	CrcUpdateChatParticipantDelete            = 0x6e5f8c22
	CrcUpdateDcOptions                        = 0x8e5e9873
	CrcInputMediaUploadedAudio                = 0x4e498cab
	CrcInputMediaAudio                        = 0x89938781
	CrcInputMediaUploadedDocument             = 0xffe76b78
	CrcInputMediaUploadedThumbDocument        = 0x41481486
	CrcInputMediaDocument                     = 0xd184e841
	CrcMessageMediaDocument                   = 0x2fda2204
	CrcMessageMediaAudio                      = 0xc6b68300
	CrcInputAudioEmpty                        = 0xd95adc84
	CrcInputAudio                             = 0x77d440ff
	CrcInputDocumentEmpty                     = 0x72f0eaae
	CrcInputDocument                          = 0x18798952
	CrcInputAudioFileLocation                 = 0x74dc404d
	CrcInputDocumentFileLocation              = 0x4e45abe9
	CrcAudioEmpty                             = 0x586988d8
	CrcAudio                                  = 0xc7ac6496
	CrcDocumentEmpty                          = 0x36f8c871
	CrcDocument                               = 0xf9a39f4f
	CrcHelpSupport                            = 0x17c6b5f6
	CrcNotifyPeer                             = 0x9fd40bd8
	CrcNotifyUsers                            = 0xb4c83b4c
	CrcNotifyChats                            = 0xc007cec3
	CrcNotifyAll                              = 0x74d07c60
	CrcUpdateUserBlocked                      = 0x80ece81a
	CrcUpdateNotifySettings                   = 0xbec268ef
	CrcAuthSentAppCode                        = 0xe325edcf
	CrcSendMessageTypingAction                = 0x16bf744e
	CrcSendMessageCancelAction                = 0xfd5ec8f5
	CrcSendMessageRecordVideoAction           = 0xa187d66f
	CrcSendMessageUploadVideoAction           = 0x92042ff7
	CrcSendMessageRecordAudioAction           = 0xd52f73f7
	CrcSendMessageUploadAudioAction           = 0xe6ac8a6f
	CrcSendMessageUploadPhotoAction           = 0x990a3c1a
	CrcSendMessageUploadDocumentAction        = 0x8faee98e
	CrcSendMessageGeoLocationAction           = 0x176f8ba1
	CrcSendMessageChooseContactAction         = 0x628cbc6f
	CrcContactFound                           = 0xea879f95
	CrcContactsFound                          = 0x0566000e
	CrcUpdateServiceNotification              = 0x382dd3e4
	CrcUserStatusRecently                     = 0xe26f42f1
	CrcUserStatusLastWeek                     = 0x07bf09fc
	CrcUserStatusLastMonth                    = 0x77ebc742
	CrcUpdatePrivacy                          = 0xee3b272a
	CrcInputPrivacyKeyStatusTimestamp         = 0x4f96cb18
	CrcPrivacyKeyStatusTimestamp              = 0xbc2eab30
	CrcInputPrivacyValueAllowContacts         = 0x0d09e07b
	CrcInputPrivacyValueAllowAll              = 0x184b35ce
	CrcInputPrivacyValueAllowUsers            = 0x131cc67f
	CrcInputPrivacyValueDisallowContacts      = 0x0ba52007
	CrcInputPrivacyValueDisallowAll           = 0xd66b66c9
	CrcInputPrivacyValueDisallowUsers         = 0x90110467
	CrcPrivacyValueAllowContacts              = 0xfffe1bac
	CrcPrivacyValueAllowAll                   = 0x65427b82
	CrcPrivacyValueAllowUsers                 = 0x4d5bbe0c
	CrcPrivacyValueDisallowContacts           = 0xf888fa1a
	CrcPrivacyValueDisallowAll                = 0x8b73e763
	CrcPrivacyValueDisallowUsers              = 0x0c7f49b7
	CrcAccountPrivacyRules                    = 0x554abb6f
	CrcAccountDaysTTL                         = 0xb8d0afdf
	CrcAccountSentChangePhoneCode             = 0xa4f58c4c
	CrcUpdateUserPhone                        = 0x12b9417b
	CrcDocumentAttributeImageSize             = 0x6c37c15c
	CrcDocumentAttributeAnimated              = 0x11b58939
	CrcDocumentAttributeSticker               = 0xfb0a5727
	CrcDocumentAttributeVideo                 = 0x5910cccb
	CrcDocumentAttributeAudio                 = 0x051448e5
	CrcDocumentAttributeFilename              = 0x15590068
	CrcMessagesStickersNotModified            = 0xf1749a22
	CrcMessagesStickers                       = 0x8a8ecd32
	CrcStickerPack                            = 0x12b299d4
	CrcMessagesAllStickersNotModified         = 0xe86602c3
	CrcMessagesAllStickers                    = 0xdcef3102
	CrcDisabledFeature                        = 0xae636f24
	CrcInvokeAfterMsg                         = 0xcb9f372d
	CrcInvokeAfterMsgs                        = 0x3dc4b4f0
	CrcAuthCheckPhone                         = 0x6fe51dfb
	CrcAuthSendCode                           = 0x768d5f4d
	CrcAuthSendCall                           = 0x03c51564
	CrcAuthSignUp                             = 0x1b067634
	CrcAuthSignIn                             = 0xbcd51581
	CrcAuthLogOut                             = 0x5717da40
	CrcAuthResetAuthorizations                = 0x9fab0d1a
	CrcAuthSendInvites                        = 0x771c1d97
	CrcAuthExportAuthorization                = 0xe5bfffcd
	CrcAuthImportAuthorization                = 0xe3ef9613
	CrcAuthBindTempAuthKey                    = 0xcdd42a05
	CrcAccountRegisterDevice                  = 0x446c712c
	CrcAccountUnregisterDevice                = 0x65c55b40
	CrcAccountUpdateNotifySettings            = 0x84be5b93
	CrcAccountGetNotifySettings               = 0x12b3ad31
	CrcAccountResetNotifySettings             = 0xdb7e1747
	CrcAccountUpdateProfile                   = 0xf0888d68
	CrcAccountUpdateStatus                    = 0x6628562c
	CrcAccountGetWallPapers                   = 0xc04cfac2
	CrcUsersGetUsers                          = 0x0d91a548
	CrcUsersGetFullUser                       = 0xca30a5b1
	CrcContactsGetStatuses                    = 0xc4a353ee
	CrcContactsGetContacts                    = 0x22c6aa08
	CrcContactsImportContacts                 = 0xda30b32d
	CrcContactsGetSuggested                   = 0xcd773428
	CrcContactsDeleteContact                  = 0x8e953744
	CrcContactsDeleteContacts                 = 0x59ab389e
	CrcContactsBlock                          = 0x332b49fc
	CrcContactsUnblock                        = 0xe54100bd
	CrcContactsGetBlocked                     = 0xf57c350f
	CrcContactsExportCard                     = 0x84e53737
	CrcContactsImportCard                     = 0x4fe196fe
	CrcMessagesGetMessages                    = 0x4222fa74
	CrcMessagesGetDialogs                     = 0xeccf1df6
	CrcMessagesGetHistory                     = 0x92a1df2f
	CrcMessagesSearch                         = 0x07e9f2ab
	CrcMessagesReadHistory                    = 0xeed884c6
	CrcMessagesDeleteHistory                  = 0xf4f8fb61
	CrcMessagesDeleteMessages                 = 0x14f2dd0a
	CrcMessagesReceivedMessages               = 0x28abcb68
	CrcMessagesSetTyping                      = 0xa3825e50
	CrcMessagesSendMessage                    = 0x4cde0aab
	CrcMessagesSendMedia                      = 0xa3c85d76
	CrcMessagesForwardMessages                = 0x514cd10f
	CrcMessagesGetChats                       = 0x3c6aa187
	CrcMessagesGetFullChat                    = 0x3b831c66
	CrcMessagesEditChatTitle                  = 0xb4bc68b5
	CrcMessagesEditChatPhoto                  = 0xd881821d
	CrcMessagesAddChatUser                    = 0x2ee9ee9e
	CrcMessagesDeleteChatUser                 = 0xc3c5cd23
	CrcMessagesCreateChat                     = 0x419d9aee
	CrcUpdatesGetState                        = 0xedd4882a
	CrcUpdatesGetDifference                   = 0x0a041495
	CrcPhotosUpdateProfilePhoto               = 0xeef579a0
	CrcPhotosUploadProfilePhoto               = 0xd50f9c88
	CrcPhotosDeletePhotos                     = 0x87cf7f2f
	CrcUploadSaveFilePart                     = 0xb304a621
	CrcUploadGetFile                          = 0xe3a6cfb5
	CrcHelpGetConfig                          = 0xc4f9186b
	CrcHelpGetNearestDc                       = 0x1fb33026
	CrcHelpGetAppUpdate                       = 0xc812ac7e
	CrcHelpSaveAppLog                         = 0x6f02f748
	CrcHelpGetInviteText                      = 0xa4a95186
	CrcPhotosGetUserPhotos                    = 0xb7ee553c
	CrcMessagesForwardMessage                 = 0x03f3f4f2
	CrcMessagesSendBroadcast                  = 0x41bb0972
	CrcGeochatsGetLocated                     = 0x7f192d8f
	CrcGeochatsGetRecents                     = 0xe1427e6f
	CrcGeochatsCheckin                        = 0x55b3e8fb
	CrcGeochatsGetFullChat                    = 0x6722dd6f
	CrcGeochatsEditChatTitle                  = 0x4c8e2273
	CrcGeochatsEditChatPhoto                  = 0x35d81a95
	CrcGeochatsSearch                         = 0xcfcdc44d
	CrcGeochatsGetHistory                     = 0xb53f7a68
	CrcGeochatsSetTyping                      = 0x08b8a729
	CrcGeochatsSendMessage                    = 0x061b0044
	CrcGeochatsSendMedia                      = 0xb8f0deff
	CrcGeochatsCreateGeoChat                  = 0x0e092e16
	CrcMessagesGetDhConfig                    = 0x26cf8950
	CrcMessagesRequestEncryption              = 0xf64daf43
	CrcMessagesAcceptEncryption               = 0x3dbc0415
	CrcMessagesDiscardEncryption              = 0xedd923c5
	CrcMessagesSetEncryptedTyping             = 0x791451ed
	CrcMessagesReadEncryptedHistory           = 0x7f4b690a
	CrcMessagesSendEncrypted                  = 0xa9776773
	CrcMessagesSendEncryptedFile              = 0x9a901b66
	CrcMessagesSendEncryptedService           = 0x32d439a4
	CrcMessagesReceivedQueue                  = 0x55a5bb66
	CrcUploadSaveBigFilePart                  = 0xde7b673d
	CrcInitConnection                         = 0x69796de9
	CrcHelpGetSupport                         = 0x9cdf08cd
	CrcAuthSendSms                            = 0x0da9f3e8
	CrcMessagesReadMessageContents            = 0x354b5bc2
	CrcAccountCheckUsername                   = 0x2714d86c
	CrcAccountUpdateUsername                  = 0x3e0bdd7c
	CrcContactsSearch                         = 0x11f812d8
	CrcAccountGetPrivacy                      = 0xdadbc950
	CrcAccountSetPrivacy                      = 0xc9f81ce8
	CrcAccountDeleteAccount                   = 0x418d4e0b
	CrcAccountGetAccountTTL                   = 0x08fc711d
	CrcAccountSetAccountTTL                   = 0x2442485e
	CrcInvokeWithLayer                        = 0xda9b0d0d
	CrcContactsResolveUsername                = 0x0bf0131c
	CrcAccountSendChangePhoneCode             = 0xa407a8f4
	CrcAccountChangePhone                     = 0x70c32edb
	CrcMessagesGetStickers                    = 0xae22e045
	CrcMessagesGetAllStickers                 = 0xaa3bc868
	CrcAccountUpdateDeviceLocked              = 0x38df3532
)

type TLBoolFalse struct {
}

type TLBoolTrue struct {
}

type TLVector struct {
}

type TLError struct {
	Code int32
	Text string
}

type TLNull struct {
}

type TLInputPeerEmpty struct {
}

type TLInputPeerSelf struct {
}

type TLInputPeerContact struct {
	UserId int32
}

type TLInputPeerForeign struct {
	UserId     int32
	AccessHash int64
}

type TLInputPeerChat struct {
	ChatId int32
}

type TLInputUserEmpty struct {
}

type TLInputUserSelf struct {
}

type TLInputUserContact struct {
	UserId int32
}

type TLInputUserForeign struct {
	UserId     int32
	AccessHash int64
}

type TLInputPhoneContact struct {
	ClientId  int64
	Phone     string
	FirstName string
	LastName  string
}

type TLInputFile struct {
	Id          int64
	Parts       int32
	Name        string
	Md5Checksum string
}

type TLInputMediaEmpty struct {
}

type TLInputMediaUploadedPhoto struct {
	File TL // InputFile
}

type TLInputMediaPhoto struct {
	Id TL // InputPhoto
}

type TLInputMediaGeoPoint struct {
	GeoPoint TL // InputGeoPoint
}

type TLInputMediaContact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
}

type TLInputMediaUploadedVideo struct {
	File     TL // InputFile
	Duration int32
	W        int32
	H        int32
	MimeType string
}

type TLInputMediaUploadedThumbVideo struct {
	File     TL // InputFile
	Thumb    TL // InputFile
	Duration int32
	W        int32
	H        int32
	MimeType string
}

type TLInputMediaVideo struct {
	Id TL // InputVideo
}

type TLInputChatPhotoEmpty struct {
}

type TLInputChatUploadedPhoto struct {
	File TL // InputFile
	Crop TL // InputPhotoCrop
}

type TLInputChatPhoto struct {
	Id   TL // InputPhoto
	Crop TL // InputPhotoCrop
}

type TLInputGeoPointEmpty struct {
}

type TLInputGeoPoint struct {
	Lat  float64
	Long float64
}

type TLInputPhotoEmpty struct {
}

type TLInputPhoto struct {
	Id         int64
	AccessHash int64
}

type TLInputVideoEmpty struct {
}

type TLInputVideo struct {
	Id         int64
	AccessHash int64
}

type TLInputFileLocation struct {
	VolumeId int64
	LocalId  int32
	Secret   int64
}

type TLInputVideoFileLocation struct {
	Id         int64
	AccessHash int64
}

type TLInputPhotoCropAuto struct {
}

type TLInputPhotoCrop struct {
	CropLeft  float64
	CropTop   float64
	CropWidth float64
}

type TLInputAppEvent struct {
	Time  float64
	_type string
	Peer  int64
	Data  string
}

type TLPeerUser struct {
	UserId int32
}

type TLPeerChat struct {
	ChatId int32
}

type TLStorageFileUnknown struct {
}

type TLStorageFileJpeg struct {
}

type TLStorageFileGif struct {
}

type TLStorageFilePng struct {
}

type TLStorageFilePdf struct {
}

type TLStorageFileMp3 struct {
}

type TLStorageFileMov struct {
}

type TLStorageFilePartial struct {
}

type TLStorageFileMp4 struct {
}

type TLStorageFileWebp struct {
}

type TLFileLocationUnavailable struct {
	VolumeId int64
	LocalId  int32
	Secret   int64
}

type TLFileLocation struct {
	DcId     int32
	VolumeId int64
	LocalId  int32
	Secret   int64
}

type TLUserEmpty struct {
	Id int32
}

type TLUserSelf struct {
	Id        int32
	FirstName string
	LastName  string
	Username  string
	Phone     string
	Photo     TL // UserProfilePhoto
	Status    TL // UserStatus
	Inactive  TL // Bool
}

type TLUserContact struct {
	Id         int32
	FirstName  string
	LastName   string
	Username   string
	AccessHash int64
	Phone      string
	Photo      TL // UserProfilePhoto
	Status     TL // UserStatus
}

type TLUserRequest struct {
	Id         int32
	FirstName  string
	LastName   string
	Username   string
	AccessHash int64
	Phone      string
	Photo      TL // UserProfilePhoto
	Status     TL // UserStatus
}

type TLUserForeign struct {
	Id         int32
	FirstName  string
	LastName   string
	Username   string
	AccessHash int64
	Photo      TL // UserProfilePhoto
	Status     TL // UserStatus
}

type TLUserDeleted struct {
	Id        int32
	FirstName string
	LastName  string
	Username  string
}

type TLUserProfilePhotoEmpty struct {
}

type TLUserProfilePhoto struct {
	PhotoId    int64
	PhotoSmall TL // FileLocation
	PhotoBig   TL // FileLocation
}

type TLUserStatusEmpty struct {
}

type TLUserStatusOnline struct {
	Expires int32
}

type TLUserStatusOffline struct {
	WasOnline int32
}

type TLChatEmpty struct {
	Id int32
}

type TLChat struct {
	Id                int32
	Title             string
	Photo             TL // ChatPhoto
	ParticipantsCount int32
	Date              int32
	Left              TL // Bool
	Version           int32
}

type TLChatForbidden struct {
	Id    int32
	Title string
	Date  int32
}

type TLChatFull struct {
	Id             int32
	Participants   TL // ChatParticipants
	ChatPhoto      TL // Photo
	NotifySettings TL // PeerNotifySettings
}

type TLChatParticipant struct {
	UserId    int32
	InviterId int32
	Date      int32
}

type TLChatParticipantsForbidden struct {
	ChatId int32
}

type TLChatParticipants struct {
	ChatId       int32
	AdminId      int32
	Participants []TL // ChatParticipant
	Version      int32
}

type TLChatPhotoEmpty struct {
}

type TLChatPhoto struct {
	PhotoSmall TL // FileLocation
	PhotoBig   TL // FileLocation
}

type TLMessageEmpty struct {
	Id int32
}

type TLMessage struct {
	Flags   int32
	Id      int32
	FromId  int32
	ToId    TL // Peer
	Date    int32
	Message string
	Media   TL // MessageMedia
}

type TLMessageForwarded struct {
	Flags     int32
	Id        int32
	FwdFromId int32
	FwdDate   int32
	FromId    int32
	ToId      TL // Peer
	Date      int32
	Message   string
	Media     TL // MessageMedia
}

type TLMessageService struct {
	Flags  int32
	Id     int32
	FromId int32
	ToId   TL // Peer
	Date   int32
	Action TL // MessageAction
}

type TLMessageMediaEmpty struct {
}

type TLMessageMediaPhoto struct {
	Photo TL // Photo
}

type TLMessageMediaVideo struct {
	Video TL // Video
}

type TLMessageMediaGeo struct {
	Geo TL // GeoPoint
}

type TLMessageMediaContact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	UserId      int32
}

type TLMessageMediaUnsupported struct {
	Bytes []byte
}

type TLMessageActionEmpty struct {
}

type TLMessageActionChatCreate struct {
	Title string
	Users []int32
}

type TLMessageActionChatEditTitle struct {
	Title string
}

type TLMessageActionChatEditPhoto struct {
	Photo TL // Photo
}

type TLMessageActionChatDeletePhoto struct {
}

type TLMessageActionChatAddUser struct {
	UserId int32
}

type TLMessageActionChatDeleteUser struct {
	UserId int32
}

type TLDialog struct {
	Peer           TL // Peer
	TopMessage     int32
	UnreadCount    int32
	NotifySettings TL // PeerNotifySettings
}

type TLPhotoEmpty struct {
	Id int64
}

type TLPhoto struct {
	Id         int64
	AccessHash int64
	UserId     int32
	Date       int32
	Caption    string
	Geo        TL   // GeoPoint
	Sizes      []TL // PhotoSize
}

type TLPhotoSizeEmpty struct {
	_type string
}

type TLPhotoSize struct {
	_type    string
	Location TL // FileLocation
	W        int32
	H        int32
	Size     int32
}

type TLPhotoCachedSize struct {
	_type    string
	Location TL // FileLocation
	W        int32
	H        int32
	Bytes    []byte
}

type TLVideoEmpty struct {
	Id int64
}

type TLVideo struct {
	Id         int64
	AccessHash int64
	UserId     int32
	Date       int32
	Caption    string
	Duration   int32
	MimeType   string
	Size       int32
	Thumb      TL // PhotoSize
	DcId       int32
	W          int32
	H          int32
}

type TLGeoPointEmpty struct {
}

type TLGeoPoint struct {
	Long float64
	Lat  float64
}

type TLAuthCheckedPhone struct {
	PhoneRegistered TL // Bool
	PhoneInvited    TL // Bool
}

type TLAuthSentCode struct {
	PhoneRegistered TL // Bool
	PhoneCodeHash   string
	SendCallTimeout int32
	IsPassword      TL // Bool
}

type TLAuthAuthorization struct {
	Expires int32
	User    TL // User
}

type TLAuthExportedAuthorization struct {
	Id    int32
	Bytes []byte
}

type TLInputNotifyPeer struct {
	Peer TL // InputPeer
}

type TLInputNotifyUsers struct {
}

type TLInputNotifyChats struct {
}

type TLInputNotifyAll struct {
}

type TLInputPeerNotifyEventsEmpty struct {
}

type TLInputPeerNotifyEventsAll struct {
}

type TLInputPeerNotifySettings struct {
	MuteUntil    int32
	Sound        string
	ShowPreviews TL // Bool
	EventsMask   int32
}

type TLPeerNotifyEventsEmpty struct {
}

type TLPeerNotifyEventsAll struct {
}

type TLPeerNotifySettingsEmpty struct {
}

type TLPeerNotifySettings struct {
	MuteUntil    int32
	Sound        string
	ShowPreviews TL // Bool
	EventsMask   int32
}

type TLWallPaper struct {
	Id    int32
	Title string
	Sizes []TL // PhotoSize
	Color int32
}

type TLUserFull struct {
	User           TL // User
	Link           TL // contacts.Link
	ProfilePhoto   TL // Photo
	NotifySettings TL // PeerNotifySettings
	Blocked        TL // Bool
	RealFirstName  string
	RealLastName   string
}

type TLContact struct {
	UserId int32
	Mutual TL // Bool
}

type TLImportedContact struct {
	UserId   int32
	ClientId int64
}

type TLContactBlocked struct {
	UserId int32
	Date   int32
}

type TLContactSuggested struct {
	UserId         int32
	MutualContacts int32
}

type TLContactStatus struct {
	UserId int32
	Status TL // UserStatus
}

type TLChatLocated struct {
	ChatId   int32
	Distance int32
}

type TLContactsForeignLinkUnknown struct {
}

type TLContactsForeignLinkRequested struct {
	HasPhone TL // Bool
}

type TLContactsForeignLinkMutual struct {
}

type TLContactsMyLinkEmpty struct {
}

type TLContactsMyLinkRequested struct {
	Contact TL // Bool
}

type TLContactsMyLinkContact struct {
}

type TLContactsLink struct {
	MyLink      TL // contacts.MyLink
	ForeignLink TL // contacts.ForeignLink
	User        TL // User
}

type TLContactsContactsNotModified struct {
}

type TLContactsContacts struct {
	Contacts []TL // Contact
	Users    []TL // User
}

type TLContactsImportedContacts struct {
	Imported      []TL // ImportedContact
	RetryContacts []int64
	Users         []TL // User
}

type TLContactsBlocked struct {
	Blocked []TL // ContactBlocked
	Users   []TL // User
}

type TLContactsBlockedSlice struct {
	Count   int32
	Blocked []TL // ContactBlocked
	Users   []TL // User
}

type TLContactsSuggested struct {
	Results []TL // ContactSuggested
	Users   []TL // User
}

type TLMessagesDialogs struct {
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TLMessagesDialogsSlice struct {
	Count    int32
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TLMessagesMessages struct {
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TLMessagesMessagesSlice struct {
	Count    int32
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TLMessagesMessageEmpty struct {
}

type TLMessagesStatedMessages struct {
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
	Pts      int32
	Seq      int32
}

type TLMessagesStatedMessage struct {
	Message TL   // Message
	Chats   []TL // Chat
	Users   []TL // User
	Pts     int32
	Seq     int32
}

type TLMessagesSentMessage struct {
	Id   int32
	Date int32
	Pts  int32
	Seq  int32
}

type TLMessagesChats struct {
	Chats []TL // Chat
	Users []TL // User
}

type TLMessagesChatFull struct {
	FullChat TL   // ChatFull
	Chats    []TL // Chat
	Users    []TL // User
}

type TLMessagesAffectedHistory struct {
	Pts    int32
	Seq    int32
	Offset int32
}

type TLInputMessagesFilterEmpty struct {
}

type TLInputMessagesFilterPhotos struct {
}

type TLInputMessagesFilterVideo struct {
}

type TLInputMessagesFilterPhotoVideo struct {
}

type TLInputMessagesFilterPhotoVideoDocuments struct {
}

type TLInputMessagesFilterDocument struct {
}

type TLInputMessagesFilterAudio struct {
}

type TLUpdateNewMessage struct {
	Message TL // Message
	Pts     int32
}

type TLUpdateMessageID struct {
	Id       int32
	RandomId int64
}

type TLUpdateReadMessages struct {
	Messages []int32
	Pts      int32
}

type TLUpdateDeleteMessages struct {
	Messages []int32
	Pts      int32
}

type TLUpdateUserTyping struct {
	UserId int32
	Action TL // SendMessageAction
}

type TLUpdateChatUserTyping struct {
	ChatId int32
	UserId int32
	Action TL // SendMessageAction
}

type TLUpdateChatParticipants struct {
	Participants TL // ChatParticipants
}

type TLUpdateUserStatus struct {
	UserId int32
	Status TL // UserStatus
}

type TLUpdateUserName struct {
	UserId    int32
	FirstName string
	LastName  string
	Username  string
}

type TLUpdateUserPhoto struct {
	UserId   int32
	Date     int32
	Photo    TL // UserProfilePhoto
	Previous TL // Bool
}

type TLUpdateContactRegistered struct {
	UserId int32
	Date   int32
}

type TLUpdateContactLink struct {
	UserId      int32
	MyLink      TL // contacts.MyLink
	ForeignLink TL // contacts.ForeignLink
}

type TLUpdateNewAuthorization struct {
	AuthKeyId int64
	Date      int32
	Device    string
	Location  string
}

type TLUpdatesState struct {
	Pts         int32
	Qts         int32
	Date        int32
	Seq         int32
	UnreadCount int32
}

type TLUpdatesDifferenceEmpty struct {
	Date int32
	Seq  int32
}

type TLUpdatesDifference struct {
	NewMessages          []TL // Message
	NewEncryptedMessages []TL // EncryptedMessage
	OtherUpdates         []TL // Update
	Chats                []TL // Chat
	Users                []TL // User
	State                TL   // updates.State
}

type TLUpdatesDifferenceSlice struct {
	NewMessages          []TL // Message
	NewEncryptedMessages []TL // EncryptedMessage
	OtherUpdates         []TL // Update
	Chats                []TL // Chat
	Users                []TL // User
	IntermediateState    TL   // updates.State
}

type TLUpdatesTooLong struct {
}

type TLUpdateShortMessage struct {
	Id      int32
	FromId  int32
	Message string
	Pts     int32
	Date    int32
	Seq     int32
}

type TLUpdateShortChatMessage struct {
	Id      int32
	FromId  int32
	ChatId  int32
	Message string
	Pts     int32
	Date    int32
	Seq     int32
}

type TLUpdateShort struct {
	Update TL // Update
	Date   int32
}

type TLUpdatesCombined struct {
	Updates  []TL // Update
	Users    []TL // User
	Chats    []TL // Chat
	Date     int32
	SeqStart int32
	Seq      int32
}

type TLUpdates struct {
	Updates []TL // Update
	Users   []TL // User
	Chats   []TL // Chat
	Date    int32
	Seq     int32
}

type TLPhotosPhotos struct {
	Photos []TL // Photo
	Users  []TL // User
}

type TLPhotosPhotosSlice struct {
	Count  int32
	Photos []TL // Photo
	Users  []TL // User
}

type TLPhotosPhoto struct {
	Photo TL   // Photo
	Users []TL // User
}

type TLUploadFile struct {
	_type TL // storage.FileType
	Mtime int32
	Bytes []byte
}

type TLDcOption struct {
	Id        int32
	Hostname  string
	IpAddress string
	Port      int32
}

type TLConfig struct {
	Date             int32
	Expires          int32
	TestMode         TL // Bool
	ThisDc           int32
	DcOptions        []TL // DcOption
	ChatBigSize      int32
	ChatSizeMax      int32
	BroadcastSizeMax int32
	DisabledFeatures []TL // DisabledFeature
}

type TLNearestDc struct {
	Country   string
	ThisDc    int32
	NearestDc int32
}

type TLHelpAppUpdate struct {
	Id       int32
	Critical TL // Bool
	Url      string
	Text     string
}

type TLHelpNoAppUpdate struct {
}

type TLHelpInviteText struct {
	Message string
}

type TLMessagesStatedMessagesLinks struct {
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
	Links    []TL // contacts.Link
	Pts      int32
	Seq      int32
}

type TLMessagesStatedMessageLink struct {
	Message TL   // Message
	Chats   []TL // Chat
	Users   []TL // User
	Links   []TL // contacts.Link
	Pts     int32
	Seq     int32
}

type TLMessagesSentMessageLink struct {
	Id    int32
	Date  int32
	Pts   int32
	Seq   int32
	Links []TL // contacts.Link
}

type TLInputGeoChat struct {
	ChatId     int32
	AccessHash int64
}

type TLInputNotifyGeoChatPeer struct {
	Peer TL // InputGeoChat
}

type TLGeoChat struct {
	Id                int32
	AccessHash        int64
	Title             string
	Address           string
	Venue             string
	Geo               TL // GeoPoint
	Photo             TL // ChatPhoto
	ParticipantsCount int32
	Date              int32
	CheckedIn         TL // Bool
	Version           int32
}

type TLGeoChatMessageEmpty struct {
	ChatId int32
	Id     int32
}

type TLGeoChatMessage struct {
	ChatId  int32
	Id      int32
	FromId  int32
	Date    int32
	Message string
	Media   TL // MessageMedia
}

type TLGeoChatMessageService struct {
	ChatId int32
	Id     int32
	FromId int32
	Date   int32
	Action TL // MessageAction
}

type TLGeochatsStatedMessage struct {
	Message TL   // GeoChatMessage
	Chats   []TL // Chat
	Users   []TL // User
	Seq     int32
}

type TLGeochatsLocated struct {
	Results  []TL // ChatLocated
	Messages []TL // GeoChatMessage
	Chats    []TL // Chat
	Users    []TL // User
}

type TLGeochatsMessages struct {
	Messages []TL // GeoChatMessage
	Chats    []TL // Chat
	Users    []TL // User
}

type TLGeochatsMessagesSlice struct {
	Count    int32
	Messages []TL // GeoChatMessage
	Chats    []TL // Chat
	Users    []TL // User
}

type TLMessageActionGeoChatCreate struct {
	Title   string
	Address string
}

type TLMessageActionGeoChatCheckin struct {
}

type TLUpdateNewGeoChatMessage struct {
	Message TL // GeoChatMessage
}

type TLWallPaperSolid struct {
	Id      int32
	Title   string
	BgColor int32
	Color   int32
}

type TLUpdateNewEncryptedMessage struct {
	Message TL // EncryptedMessage
	Qts     int32
}

type TLUpdateEncryptedChatTyping struct {
	ChatId int32
}

type TLUpdateEncryption struct {
	Chat TL // EncryptedChat
	Date int32
}

type TLUpdateEncryptedMessagesRead struct {
	ChatId  int32
	MaxDate int32
	Date    int32
}

type TLEncryptedChatEmpty struct {
	Id int32
}

type TLEncryptedChatWaiting struct {
	Id            int32
	AccessHash    int64
	Date          int32
	AdminId       int32
	ParticipantId int32
}

type TLEncryptedChatRequested struct {
	Id            int32
	AccessHash    int64
	Date          int32
	AdminId       int32
	ParticipantId int32
	GA            []byte
}

type TLEncryptedChat struct {
	Id             int32
	AccessHash     int64
	Date           int32
	AdminId        int32
	ParticipantId  int32
	GAOrB          []byte
	KeyFingerprint int64
}

type TLEncryptedChatDiscarded struct {
	Id int32
}

type TLInputEncryptedChat struct {
	ChatId     int32
	AccessHash int64
}

type TLEncryptedFileEmpty struct {
}

type TLEncryptedFile struct {
	Id             int64
	AccessHash     int64
	Size           int32
	DcId           int32
	KeyFingerprint int32
}

type TLInputEncryptedFileEmpty struct {
}

type TLInputEncryptedFileUploaded struct {
	Id             int64
	Parts          int32
	Md5Checksum    string
	KeyFingerprint int32
}

type TLInputEncryptedFile struct {
	Id         int64
	AccessHash int64
}

type TLInputEncryptedFileLocation struct {
	Id         int64
	AccessHash int64
}

type TLEncryptedMessage struct {
	RandomId int64
	ChatId   int32
	Date     int32
	Bytes    []byte
	File     TL // EncryptedFile
}

type TLEncryptedMessageService struct {
	RandomId int64
	ChatId   int32
	Date     int32
	Bytes    []byte
}

type TLMessagesDhConfigNotModified struct {
	Random []byte
}

type TLMessagesDhConfig struct {
	G       int32
	P       []byte
	Version int32
	Random  []byte
}

type TLMessagesSentEncryptedMessage struct {
	Date int32
}

type TLMessagesSentEncryptedFile struct {
	Date int32
	File TL // EncryptedFile
}

type TLInputFileBig struct {
	Id    int64
	Parts int32
	Name  string
}

type TLInputEncryptedFileBigUploaded struct {
	Id             int64
	Parts          int32
	KeyFingerprint int32
}

type TLUpdateChatParticipantAdd struct {
	ChatId    int32
	UserId    int32
	InviterId int32
	Version   int32
}

type TLUpdateChatParticipantDelete struct {
	ChatId  int32
	UserId  int32
	Version int32
}

type TLUpdateDcOptions struct {
	DcOptions []TL // DcOption
}

type TLInputMediaUploadedAudio struct {
	File     TL // InputFile
	Duration int32
	MimeType string
}

type TLInputMediaAudio struct {
	Id TL // InputAudio
}

type TLInputMediaUploadedDocument struct {
	File       TL // InputFile
	MimeType   string
	Attributes []TL // DocumentAttribute
}

type TLInputMediaUploadedThumbDocument struct {
	File       TL // InputFile
	Thumb      TL // InputFile
	MimeType   string
	Attributes []TL // DocumentAttribute
}

type TLInputMediaDocument struct {
	Id TL // InputDocument
}

type TLMessageMediaDocument struct {
	Document TL // Document
}

type TLMessageMediaAudio struct {
	Audio TL // Audio
}

type TLInputAudioEmpty struct {
}

type TLInputAudio struct {
	Id         int64
	AccessHash int64
}

type TLInputDocumentEmpty struct {
}

type TLInputDocument struct {
	Id         int64
	AccessHash int64
}

type TLInputAudioFileLocation struct {
	Id         int64
	AccessHash int64
}

type TLInputDocumentFileLocation struct {
	Id         int64
	AccessHash int64
}

type TLAudioEmpty struct {
	Id int64
}

type TLAudio struct {
	Id         int64
	AccessHash int64
	UserId     int32
	Date       int32
	Duration   int32
	MimeType   string
	Size       int32
	DcId       int32
}

type TLDocumentEmpty struct {
	Id int64
}

type TLDocument struct {
	Id         int64
	AccessHash int64
	Date       int32
	MimeType   string
	Size       int32
	Thumb      TL // PhotoSize
	DcId       int32
	Attributes []TL // DocumentAttribute
}

type TLHelpSupport struct {
	PhoneNumber string
	User        TL // User
}

type TLNotifyPeer struct {
	Peer TL // Peer
}

type TLNotifyUsers struct {
}

type TLNotifyChats struct {
}

type TLNotifyAll struct {
}

type TLUpdateUserBlocked struct {
	UserId  int32
	Blocked TL // Bool
}

type TLUpdateNotifySettings struct {
	Peer           TL // NotifyPeer
	NotifySettings TL // PeerNotifySettings
}

type TLAuthSentAppCode struct {
	PhoneRegistered TL // Bool
	PhoneCodeHash   string
	SendCallTimeout int32
	IsPassword      TL // Bool
}

type TLSendMessageTypingAction struct {
}

type TLSendMessageCancelAction struct {
}

type TLSendMessageRecordVideoAction struct {
}

type TLSendMessageUploadVideoAction struct {
}

type TLSendMessageRecordAudioAction struct {
}

type TLSendMessageUploadAudioAction struct {
}

type TLSendMessageUploadPhotoAction struct {
}

type TLSendMessageUploadDocumentAction struct {
}

type TLSendMessageGeoLocationAction struct {
}

type TLSendMessageChooseContactAction struct {
}

type TLContactFound struct {
	UserId int32
}

type TLContactsFound struct {
	Results []TL // ContactFound
	Users   []TL // User
}

type TLUpdateServiceNotification struct {
	_type   string
	Message string
	Media   TL // MessageMedia
	Popup   TL // Bool
}

type TLUserStatusRecently struct {
}

type TLUserStatusLastWeek struct {
}

type TLUserStatusLastMonth struct {
}

type TLUpdatePrivacy struct {
	Key   TL   // PrivacyKey
	Rules []TL // PrivacyRule
}

type TLInputPrivacyKeyStatusTimestamp struct {
}

type TLPrivacyKeyStatusTimestamp struct {
}

type TLInputPrivacyValueAllowContacts struct {
}

type TLInputPrivacyValueAllowAll struct {
}

type TLInputPrivacyValueAllowUsers struct {
	Users []TL // InputUser
}

type TLInputPrivacyValueDisallowContacts struct {
}

type TLInputPrivacyValueDisallowAll struct {
}

type TLInputPrivacyValueDisallowUsers struct {
	Users []TL // InputUser
}

type TLPrivacyValueAllowContacts struct {
}

type TLPrivacyValueAllowAll struct {
}

type TLPrivacyValueAllowUsers struct {
	Users []int32
}

type TLPrivacyValueDisallowContacts struct {
}

type TLPrivacyValueDisallowAll struct {
}

type TLPrivacyValueDisallowUsers struct {
	Users []int32
}

type TLAccountPrivacyRules struct {
	Rules []TL // PrivacyRule
	Users []TL // User
}

type TLAccountDaysTTL struct {
	Days int32
}

type TLAccountSentChangePhoneCode struct {
	PhoneCodeHash   string
	SendCallTimeout int32
}

type TLUpdateUserPhone struct {
	UserId int32
	Phone  string
}

type TLDocumentAttributeImageSize struct {
	W int32
	H int32
}

type TLDocumentAttributeAnimated struct {
}

type TLDocumentAttributeSticker struct {
}

type TLDocumentAttributeVideo struct {
	Duration int32
	W        int32
	H        int32
}

type TLDocumentAttributeAudio struct {
	Duration int32
}

type TLDocumentAttributeFilename struct {
	FileName string
}

type TLMessagesStickersNotModified struct {
}

type TLMessagesStickers struct {
	Hash     string
	Stickers []TL // Document
}

type TLStickerPack struct {
	Emoticon  string
	Documents []int64
}

type TLMessagesAllStickersNotModified struct {
}

type TLMessagesAllStickers struct {
	Hash      string
	Packs     []TL // StickerPack
	Documents []TL // Document
}

type TLDisabledFeature struct {
	Feature     string
	Description string
}

type TLInvokeAfterMsg struct {
	MsgId int64
	Query TL
}

type TLInvokeAfterMsgs struct {
	MsgIds []int64
	Query  TL
}

type TLAuthCheckPhone struct {
	PhoneNumber string
}

type TLAuthSendCode struct {
	PhoneNumber string
	SmsType     int32
	ApiId       int32
	ApiHash     string
	LangCode    string
}

type TLAuthSendCall struct {
	PhoneNumber   string
	PhoneCodeHash string
}

type TLAuthSignUp struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
	FirstName     string
	LastName      string
}

type TLAuthSignIn struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

type TLAuthLogOut struct {
}

type TLAuthResetAuthorizations struct {
}

type TLAuthSendInvites struct {
	PhoneNumbers []string
	Message      string
}

type TLAuthExportAuthorization struct {
	DcId int32
}

type TLAuthImportAuthorization struct {
	Id    int32
	Bytes []byte
}

type TLAuthBindTempAuthKey struct {
	PermAuthKeyId    int64
	Nonce            int64
	ExpiresAt        int32
	EncryptedMessage []byte
}

type TLAccountRegisterDevice struct {
	TokenType     int32
	Token         string
	DeviceModel   string
	SystemVersion string
	AppVersion    string
	AppSandbox    TL // Bool
	LangCode      string
}

type TLAccountUnregisterDevice struct {
	TokenType int32
	Token     string
}

type TLAccountUpdateNotifySettings struct {
	Peer     TL // InputNotifyPeer
	Settings TL // InputPeerNotifySettings
}

type TLAccountGetNotifySettings struct {
	Peer TL // InputNotifyPeer
}

type TLAccountResetNotifySettings struct {
}

type TLAccountUpdateProfile struct {
	FirstName string
	LastName  string
}

type TLAccountUpdateStatus struct {
	Offline TL // Bool
}

type TLAccountGetWallPapers struct {
}

type TLUsersGetUsers struct {
	Id []TL // InputUser
}

type TLUsersGetFullUser struct {
	Id TL // InputUser
}

type TLContactsGetStatuses struct {
}

type TLContactsGetContacts struct {
	Hash string
}

type TLContactsImportContacts struct {
	Contacts []TL // InputContact
	Replace  TL   // Bool
}

type TLContactsGetSuggested struct {
	Limit int32
}

type TLContactsDeleteContact struct {
	Id TL // InputUser
}

type TLContactsDeleteContacts struct {
	Id []TL // InputUser
}

type TLContactsBlock struct {
	Id TL // InputUser
}

type TLContactsUnblock struct {
	Id TL // InputUser
}

type TLContactsGetBlocked struct {
	Offset int32
	Limit  int32
}

type TLContactsExportCard struct {
}

type TLContactsImportCard struct {
	ExportCard []int32
}

type TLMessagesGetMessages struct {
	Id []int32
}

type TLMessagesGetDialogs struct {
	Offset int32
	MaxId  int32
	Limit  int32
}

type TLMessagesGetHistory struct {
	Peer   TL // InputPeer
	Offset int32
	MaxId  int32
	Limit  int32
}

type TLMessagesSearch struct {
	Peer    TL // InputPeer
	Q       string
	Filter  TL // MessagesFilter
	MinDate int32
	MaxDate int32
	Offset  int32
	MaxId   int32
	Limit   int32
}

type TLMessagesReadHistory struct {
	Peer         TL // InputPeer
	MaxId        int32
	Offset       int32
	ReadContents TL // Bool
}

type TLMessagesDeleteHistory struct {
	Peer   TL // InputPeer
	Offset int32
}

type TLMessagesDeleteMessages struct {
	Id []int32
}

type TLMessagesReceivedMessages struct {
	MaxId int32
}

type TLMessagesSetTyping struct {
	Peer   TL // InputPeer
	Action TL // SendMessageAction
}

type TLMessagesSendMessage struct {
	Peer     TL // InputPeer
	Message  string
	RandomId int64
}

type TLMessagesSendMedia struct {
	Peer     TL // InputPeer
	Media    TL // InputMedia
	RandomId int64
}

type TLMessagesForwardMessages struct {
	Peer TL // InputPeer
	Id   []int32
}

type TLMessagesGetChats struct {
	Id []int32
}

type TLMessagesGetFullChat struct {
	ChatId int32
}

type TLMessagesEditChatTitle struct {
	ChatId int32
	Title  string
}

type TLMessagesEditChatPhoto struct {
	ChatId int32
	Photo  TL // InputChatPhoto
}

type TLMessagesAddChatUser struct {
	ChatId   int32
	UserId   TL // InputUser
	FwdLimit int32
}

type TLMessagesDeleteChatUser struct {
	ChatId int32
	UserId TL // InputUser
}

type TLMessagesCreateChat struct {
	Users []TL // InputUser
	Title string
}

type TLUpdatesGetState struct {
}

type TLUpdatesGetDifference struct {
	Pts  int32
	Date int32
	Qts  int32
}

type TLPhotosUpdateProfilePhoto struct {
	Id   TL // InputPhoto
	Crop TL // InputPhotoCrop
}

type TLPhotosUploadProfilePhoto struct {
	File     TL // InputFile
	Caption  string
	GeoPoint TL // InputGeoPoint
	Crop     TL // InputPhotoCrop
}

type TLPhotosDeletePhotos struct {
	Id []TL // InputPhoto
}

type TLUploadSaveFilePart struct {
	FileId   int64
	FilePart int32
	Bytes    []byte
}

type TLUploadGetFile struct {
	Location TL // InputFileLocation
	Offset   int32
	Limit    int32
}

type TLHelpGetConfig struct {
}

type TLHelpGetNearestDc struct {
}

type TLHelpGetAppUpdate struct {
	DeviceModel   string
	SystemVersion string
	AppVersion    string
	LangCode      string
}

type TLHelpSaveAppLog struct {
	Events []TL // InputAppEvent
}

type TLHelpGetInviteText struct {
	LangCode string
}

type TLPhotosGetUserPhotos struct {
	UserId TL // InputUser
	Offset int32
	MaxId  int32
	Limit  int32
}

type TLMessagesForwardMessage struct {
	Peer     TL // InputPeer
	Id       int32
	RandomId int64
}

type TLMessagesSendBroadcast struct {
	Contacts []TL // InputUser
	Message  string
	Media    TL // InputMedia
}

type TLGeochatsGetLocated struct {
	GeoPoint TL // InputGeoPoint
	Radius   int32
	Limit    int32
}

type TLGeochatsGetRecents struct {
	Offset int32
	Limit  int32
}

type TLGeochatsCheckin struct {
	Peer TL // InputGeoChat
}

type TLGeochatsGetFullChat struct {
	Peer TL // InputGeoChat
}

type TLGeochatsEditChatTitle struct {
	Peer    TL // InputGeoChat
	Title   string
	Address string
}

type TLGeochatsEditChatPhoto struct {
	Peer  TL // InputGeoChat
	Photo TL // InputChatPhoto
}

type TLGeochatsSearch struct {
	Peer    TL // InputGeoChat
	Q       string
	Filter  TL // MessagesFilter
	MinDate int32
	MaxDate int32
	Offset  int32
	MaxId   int32
	Limit   int32
}

type TLGeochatsGetHistory struct {
	Peer   TL // InputGeoChat
	Offset int32
	MaxId  int32
	Limit  int32
}

type TLGeochatsSetTyping struct {
	Peer   TL // InputGeoChat
	Typing TL // Bool
}

type TLGeochatsSendMessage struct {
	Peer     TL // InputGeoChat
	Message  string
	RandomId int64
}

type TLGeochatsSendMedia struct {
	Peer     TL // InputGeoChat
	Media    TL // InputMedia
	RandomId int64
}

type TLGeochatsCreateGeoChat struct {
	Title    string
	GeoPoint TL // InputGeoPoint
	Address  string
	Venue    string
}

type TLMessagesGetDhConfig struct {
	Version      int32
	RandomLength int32
}

type TLMessagesRequestEncryption struct {
	UserId   TL // InputUser
	RandomId int32
	GA       []byte
}

type TLMessagesAcceptEncryption struct {
	Peer           TL // InputEncryptedChat
	GB             []byte
	KeyFingerprint int64
}

type TLMessagesDiscardEncryption struct {
	ChatId int32
}

type TLMessagesSetEncryptedTyping struct {
	Peer   TL // InputEncryptedChat
	Typing TL // Bool
}

type TLMessagesReadEncryptedHistory struct {
	Peer    TL // InputEncryptedChat
	MaxDate int32
}

type TLMessagesSendEncrypted struct {
	Peer     TL // InputEncryptedChat
	RandomId int64
	Data     []byte
}

type TLMessagesSendEncryptedFile struct {
	Peer     TL // InputEncryptedChat
	RandomId int64
	Data     []byte
	File     TL // InputEncryptedFile
}

type TLMessagesSendEncryptedService struct {
	Peer     TL // InputEncryptedChat
	RandomId int64
	Data     []byte
}

type TLMessagesReceivedQueue struct {
	MaxQts int32
}

type TLUploadSaveBigFilePart struct {
	FileId         int64
	FilePart       int32
	FileTotalParts int32
	Bytes          []byte
}

type TLInitConnection struct {
	ApiId         int32
	DeviceModel   string
	SystemVersion string
	AppVersion    string
	LangCode      string
	Query         TL
}

type TLHelpGetSupport struct {
}

type TLAuthSendSms struct {
	PhoneNumber   string
	PhoneCodeHash string
}

type TLMessagesReadMessageContents struct {
	Id []int32
}

type TLAccountCheckUsername struct {
	Username string
}

type TLAccountUpdateUsername struct {
	Username string
}

type TLContactsSearch struct {
	Q     string
	Limit int32
}

type TLAccountGetPrivacy struct {
	Key TL // InputPrivacyKey
}

type TLAccountSetPrivacy struct {
	Key   TL   // InputPrivacyKey
	Rules []TL // InputPrivacyRule
}

type TLAccountDeleteAccount struct {
	Reason string
}

type TLAccountGetAccountTTL struct {
}

type TLAccountSetAccountTTL struct {
	Ttl TL // AccountDaysTTL
}

type TLInvokeWithLayer struct {
	Layer int32
	Query TL
}

type TLContactsResolveUsername struct {
	Username string
}

type TLAccountSendChangePhoneCode struct {
	PhoneNumber string
}

type TLAccountChangePhone struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

type TLMessagesGetStickers struct {
	Emoticon string
	Hash     string
}

type TLMessagesGetAllStickers struct {
	Hash string
}

type TLAccountUpdateDeviceLocked struct {
	Period int32
}

func (e TLBoolFalse) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcBoolFalse)
	return x.buf
}

func (e TLBoolTrue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcBoolTrue)
	return x.buf
}

func (e TLVector) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcVector)
	return x.buf
}

func (e TLError) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcError)
	x.Int(e.Code)
	x.String(e.Text)
	return x.buf
}

func (e TLNull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcNull)
	return x.buf
}

func (e TLInputPeerEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPeerEmpty)
	return x.buf
}

func (e TLInputPeerSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPeerSelf)
	return x.buf
}

func (e TLInputPeerContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPeerContact)
	x.Int(e.UserId)
	return x.buf
}

func (e TLInputPeerForeign) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPeerForeign)
	x.Int(e.UserId)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLInputPeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPeerChat)
	x.Int(e.ChatId)
	return x.buf
}

func (e TLInputUserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputUserEmpty)
	return x.buf
}

func (e TLInputUserSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputUserSelf)
	return x.buf
}

func (e TLInputUserContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputUserContact)
	x.Int(e.UserId)
	return x.buf
}

func (e TLInputUserForeign) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputUserForeign)
	x.Int(e.UserId)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLInputPhoneContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPhoneContact)
	x.Long(e.ClientId)
	x.String(e.Phone)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}

func (e TLInputFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputFile)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Name)
	x.String(e.Md5Checksum)
	return x.buf
}

func (e TLInputMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaEmpty)
	return x.buf
}

func (e TLInputMediaUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaUploadedPhoto)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TLInputMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaPhoto)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TLInputMediaGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaGeoPoint)
	x.Bytes(e.GeoPoint.encode())
	return x.buf
}

func (e TLInputMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaContact)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}

func (e TLInputMediaUploadedVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaUploadedVideo)
	x.Bytes(e.File.encode())
	x.Int(e.Duration)
	x.Int(e.W)
	x.Int(e.H)
	x.String(e.MimeType)
	return x.buf
}

func (e TLInputMediaUploadedThumbVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaUploadedThumbVideo)
	x.Bytes(e.File.encode())
	x.Bytes(e.Thumb.encode())
	x.Int(e.Duration)
	x.Int(e.W)
	x.Int(e.H)
	x.String(e.MimeType)
	return x.buf
}

func (e TLInputMediaVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaVideo)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TLInputChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputChatPhotoEmpty)
	return x.buf
}

func (e TLInputChatUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputChatUploadedPhoto)
	x.Bytes(e.File.encode())
	x.Bytes(e.Crop.encode())
	return x.buf
}

func (e TLInputChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputChatPhoto)
	x.Bytes(e.Id.encode())
	x.Bytes(e.Crop.encode())
	return x.buf
}

func (e TLInputGeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputGeoPointEmpty)
	return x.buf
}

func (e TLInputGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputGeoPoint)
	x.Double(e.Lat)
	x.Double(e.Long)
	return x.buf
}

func (e TLInputPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPhotoEmpty)
	return x.buf
}

func (e TLInputPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPhoto)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLInputVideoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputVideoEmpty)
	return x.buf
}

func (e TLInputVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputVideo)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLInputFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputFileLocation)
	x.Long(e.VolumeId)
	x.Int(e.LocalId)
	x.Long(e.Secret)
	return x.buf
}

func (e TLInputVideoFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputVideoFileLocation)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLInputPhotoCropAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPhotoCropAuto)
	return x.buf
}

func (e TLInputPhotoCrop) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPhotoCrop)
	x.Double(e.CropLeft)
	x.Double(e.CropTop)
	x.Double(e.CropWidth)
	return x.buf
}

func (e TLInputAppEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputAppEvent)
	x.Double(e.Time)
	x.String(e._type)
	x.Long(e.Peer)
	x.String(e.Data)
	return x.buf
}

func (e TLPeerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPeerUser)
	x.Int(e.UserId)
	return x.buf
}

func (e TLPeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPeerChat)
	x.Int(e.ChatId)
	return x.buf
}

func (e TLStorageFileUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStorageFileUnknown)
	return x.buf
}

func (e TLStorageFileJpeg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStorageFileJpeg)
	return x.buf
}

func (e TLStorageFileGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStorageFileGif)
	return x.buf
}

func (e TLStorageFilePng) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStorageFilePng)
	return x.buf
}

func (e TLStorageFilePdf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStorageFilePdf)
	return x.buf
}

func (e TLStorageFileMp3) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStorageFileMp3)
	return x.buf
}

func (e TLStorageFileMov) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStorageFileMov)
	return x.buf
}

func (e TLStorageFilePartial) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStorageFilePartial)
	return x.buf
}

func (e TLStorageFileMp4) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStorageFileMp4)
	return x.buf
}

func (e TLStorageFileWebp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStorageFileWebp)
	return x.buf
}

func (e TLFileLocationUnavailable) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcFileLocationUnavailable)
	x.Long(e.VolumeId)
	x.Int(e.LocalId)
	x.Long(e.Secret)
	return x.buf
}

func (e TLFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcFileLocation)
	x.Int(e.DcId)
	x.Long(e.VolumeId)
	x.Int(e.LocalId)
	x.Long(e.Secret)
	return x.buf
}

func (e TLUserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserEmpty)
	x.Int(e.Id)
	return x.buf
}

func (e TLUserSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserSelf)
	x.Int(e.Id)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Username)
	x.String(e.Phone)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Status.encode())
	x.Bytes(e.Inactive.encode())
	return x.buf
}

func (e TLUserContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserContact)
	x.Int(e.Id)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Username)
	x.Long(e.AccessHash)
	x.String(e.Phone)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e TLUserRequest) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserRequest)
	x.Int(e.Id)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Username)
	x.Long(e.AccessHash)
	x.String(e.Phone)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e TLUserForeign) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserForeign)
	x.Int(e.Id)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Username)
	x.Long(e.AccessHash)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e TLUserDeleted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserDeleted)
	x.Int(e.Id)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Username)
	return x.buf
}

func (e TLUserProfilePhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserProfilePhotoEmpty)
	return x.buf
}

func (e TLUserProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserProfilePhoto)
	x.Long(e.PhotoId)
	x.Bytes(e.PhotoSmall.encode())
	x.Bytes(e.PhotoBig.encode())
	return x.buf
}

func (e TLUserStatusEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserStatusEmpty)
	return x.buf
}

func (e TLUserStatusOnline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserStatusOnline)
	x.Int(e.Expires)
	return x.buf
}

func (e TLUserStatusOffline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserStatusOffline)
	x.Int(e.WasOnline)
	return x.buf
}

func (e TLChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcChatEmpty)
	x.Int(e.Id)
	return x.buf
}

func (e TLChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcChat)
	x.Int(e.Id)
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.ParticipantsCount)
	x.Int(e.Date)
	x.Bytes(e.Left.encode())
	x.Int(e.Version)
	return x.buf
}

func (e TLChatForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcChatForbidden)
	x.Int(e.Id)
	x.String(e.Title)
	x.Int(e.Date)
	return x.buf
}

func (e TLChatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcChatFull)
	x.Int(e.Id)
	x.Bytes(e.Participants.encode())
	x.Bytes(e.ChatPhoto.encode())
	x.Bytes(e.NotifySettings.encode())
	return x.buf
}

func (e TLChatParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcChatParticipant)
	x.Int(e.UserId)
	x.Int(e.InviterId)
	x.Int(e.Date)
	return x.buf
}

func (e TLChatParticipantsForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcChatParticipantsForbidden)
	x.Int(e.ChatId)
	return x.buf
}

func (e TLChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcChatParticipants)
	x.Int(e.ChatId)
	x.Int(e.AdminId)
	x.Vector(e.Participants)
	x.Int(e.Version)
	return x.buf
}

func (e TLChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcChatPhotoEmpty)
	return x.buf
}

func (e TLChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcChatPhoto)
	x.Bytes(e.PhotoSmall.encode())
	x.Bytes(e.PhotoBig.encode())
	return x.buf
}

func (e TLMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageEmpty)
	x.Int(e.Id)
	return x.buf
}

func (e TLMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessage)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.FromId)
	x.Bytes(e.ToId.encode())
	x.Int(e.Date)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	return x.buf
}

func (e TLMessageForwarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageForwarded)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.FwdFromId)
	x.Int(e.FwdDate)
	x.Int(e.FromId)
	x.Bytes(e.ToId.encode())
	x.Int(e.Date)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	return x.buf
}

func (e TLMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageService)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.FromId)
	x.Bytes(e.ToId.encode())
	x.Int(e.Date)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TLMessageMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageMediaEmpty)
	return x.buf
}

func (e TLMessageMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageMediaPhoto)
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e TLMessageMediaVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageMediaVideo)
	x.Bytes(e.Video.encode())
	return x.buf
}

func (e TLMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageMediaGeo)
	x.Bytes(e.Geo.encode())
	return x.buf
}

func (e TLMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageMediaContact)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.Int(e.UserId)
	return x.buf
}

func (e TLMessageMediaUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageMediaUnsupported)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TLMessageActionEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageActionEmpty)
	return x.buf
}

func (e TLMessageActionChatCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageActionChatCreate)
	x.String(e.Title)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TLMessageActionChatEditTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageActionChatEditTitle)
	x.String(e.Title)
	return x.buf
}

func (e TLMessageActionChatEditPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageActionChatEditPhoto)
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e TLMessageActionChatDeletePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageActionChatDeletePhoto)
	return x.buf
}

func (e TLMessageActionChatAddUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageActionChatAddUser)
	x.Int(e.UserId)
	return x.buf
}

func (e TLMessageActionChatDeleteUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageActionChatDeleteUser)
	x.Int(e.UserId)
	return x.buf
}

func (e TLDialog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDialog)
	x.Bytes(e.Peer.encode())
	x.Int(e.TopMessage)
	x.Int(e.UnreadCount)
	x.Bytes(e.NotifySettings.encode())
	return x.buf
}

func (e TLPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotoEmpty)
	x.Long(e.Id)
	return x.buf
}

func (e TLPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhoto)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.UserId)
	x.Int(e.Date)
	x.String(e.Caption)
	x.Bytes(e.Geo.encode())
	x.Vector(e.Sizes)
	return x.buf
}

func (e TLPhotoSizeEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotoSizeEmpty)
	x.String(e._type)
	return x.buf
}

func (e TLPhotoSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotoSize)
	x.String(e._type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Size)
	return x.buf
}

func (e TLPhotoCachedSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotoCachedSize)
	x.String(e._type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TLVideoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcVideoEmpty)
	x.Long(e.Id)
	return x.buf
}

func (e TLVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcVideo)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.UserId)
	x.Int(e.Date)
	x.String(e.Caption)
	x.Int(e.Duration)
	x.String(e.MimeType)
	x.Int(e.Size)
	x.Bytes(e.Thumb.encode())
	x.Int(e.DcId)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TLGeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeoPointEmpty)
	return x.buf
}

func (e TLGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeoPoint)
	x.Double(e.Long)
	x.Double(e.Lat)
	return x.buf
}

func (e TLAuthCheckedPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthCheckedPhone)
	x.Bytes(e.PhoneRegistered.encode())
	x.Bytes(e.PhoneInvited.encode())
	return x.buf
}

func (e TLAuthSentCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthSentCode)
	x.Bytes(e.PhoneRegistered.encode())
	x.String(e.PhoneCodeHash)
	x.Int(e.SendCallTimeout)
	x.Bytes(e.IsPassword.encode())
	return x.buf
}

func (e TLAuthAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthAuthorization)
	x.Int(e.Expires)
	x.Bytes(e.User.encode())
	return x.buf
}

func (e TLAuthExportedAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthExportedAuthorization)
	x.Int(e.Id)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TLInputNotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputNotifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TLInputNotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputNotifyUsers)
	return x.buf
}

func (e TLInputNotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputNotifyChats)
	return x.buf
}

func (e TLInputNotifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputNotifyAll)
	return x.buf
}

func (e TLInputPeerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPeerNotifyEventsEmpty)
	return x.buf
}

func (e TLInputPeerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPeerNotifyEventsAll)
	return x.buf
}

func (e TLInputPeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPeerNotifySettings)
	x.Int(e.MuteUntil)
	x.String(e.Sound)
	x.Bytes(e.ShowPreviews.encode())
	x.Int(e.EventsMask)
	return x.buf
}

func (e TLPeerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPeerNotifyEventsEmpty)
	return x.buf
}

func (e TLPeerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPeerNotifyEventsAll)
	return x.buf
}

func (e TLPeerNotifySettingsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPeerNotifySettingsEmpty)
	return x.buf
}

func (e TLPeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPeerNotifySettings)
	x.Int(e.MuteUntil)
	x.String(e.Sound)
	x.Bytes(e.ShowPreviews.encode())
	x.Int(e.EventsMask)
	return x.buf
}

func (e TLWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcWallPaper)
	x.Int(e.Id)
	x.String(e.Title)
	x.Vector(e.Sizes)
	x.Int(e.Color)
	return x.buf
}

func (e TLUserFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserFull)
	x.Bytes(e.User.encode())
	x.Bytes(e.Link.encode())
	x.Bytes(e.ProfilePhoto.encode())
	x.Bytes(e.NotifySettings.encode())
	x.Bytes(e.Blocked.encode())
	x.String(e.RealFirstName)
	x.String(e.RealLastName)
	return x.buf
}

func (e TLContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContact)
	x.Int(e.UserId)
	x.Bytes(e.Mutual.encode())
	return x.buf
}

func (e TLImportedContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcImportedContact)
	x.Int(e.UserId)
	x.Long(e.ClientId)
	return x.buf
}

func (e TLContactBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactBlocked)
	x.Int(e.UserId)
	x.Int(e.Date)
	return x.buf
}

func (e TLContactSuggested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactSuggested)
	x.Int(e.UserId)
	x.Int(e.MutualContacts)
	return x.buf
}

func (e TLContactStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactStatus)
	x.Int(e.UserId)
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e TLChatLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcChatLocated)
	x.Int(e.ChatId)
	x.Int(e.Distance)
	return x.buf
}

func (e TLContactsForeignLinkUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsForeignLinkUnknown)
	return x.buf
}

func (e TLContactsForeignLinkRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsForeignLinkRequested)
	x.Bytes(e.HasPhone.encode())
	return x.buf
}

func (e TLContactsForeignLinkMutual) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsForeignLinkMutual)
	return x.buf
}

func (e TLContactsMyLinkEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsMyLinkEmpty)
	return x.buf
}

func (e TLContactsMyLinkRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsMyLinkRequested)
	x.Bytes(e.Contact.encode())
	return x.buf
}

func (e TLContactsMyLinkContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsMyLinkContact)
	return x.buf
}

func (e TLContactsLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsLink)
	x.Bytes(e.MyLink.encode())
	x.Bytes(e.ForeignLink.encode())
	x.Bytes(e.User.encode())
	return x.buf
}

func (e TLContactsContactsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsContactsNotModified)
	return x.buf
}

func (e TLContactsContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsContacts)
	x.Vector(e.Contacts)
	x.Vector(e.Users)
	return x.buf
}

func (e TLContactsImportedContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsImportedContacts)
	x.Vector(e.Imported)
	x.VectorLong(e.RetryContacts)
	x.Vector(e.Users)
	return x.buf
}

func (e TLContactsBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsBlocked)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.buf
}

func (e TLContactsBlockedSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsBlockedSlice)
	x.Int(e.Count)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.buf
}

func (e TLContactsSuggested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsSuggested)
	x.Vector(e.Results)
	x.Vector(e.Users)
	return x.buf
}

func (e TLMessagesDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesDialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TLMessagesDialogsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesDialogsSlice)
	x.Int(e.Count)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TLMessagesMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesMessages)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TLMessagesMessagesSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesMessagesSlice)
	x.Int(e.Count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TLMessagesMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesMessageEmpty)
	return x.buf
}

func (e TLMessagesStatedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesStatedMessages)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Int(e.Pts)
	x.Int(e.Seq)
	return x.buf
}

func (e TLMessagesStatedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesStatedMessage)
	x.Bytes(e.Message.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Int(e.Pts)
	x.Int(e.Seq)
	return x.buf
}

func (e TLMessagesSentMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSentMessage)
	x.Int(e.Id)
	x.Int(e.Date)
	x.Int(e.Pts)
	x.Int(e.Seq)
	return x.buf
}

func (e TLMessagesChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesChats)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TLMessagesChatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesChatFull)
	x.Bytes(e.FullChat.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TLMessagesAffectedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesAffectedHistory)
	x.Int(e.Pts)
	x.Int(e.Seq)
	x.Int(e.Offset)
	return x.buf
}

func (e TLInputMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMessagesFilterEmpty)
	return x.buf
}

func (e TLInputMessagesFilterPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMessagesFilterPhotos)
	return x.buf
}

func (e TLInputMessagesFilterVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMessagesFilterVideo)
	return x.buf
}

func (e TLInputMessagesFilterPhotoVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMessagesFilterPhotoVideo)
	return x.buf
}

func (e TLInputMessagesFilterPhotoVideoDocuments) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMessagesFilterPhotoVideoDocuments)
	return x.buf
}

func (e TLInputMessagesFilterDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMessagesFilterDocument)
	return x.buf
}

func (e TLInputMessagesFilterAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMessagesFilterAudio)
	return x.buf
}

func (e TLUpdateNewMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateNewMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	return x.buf
}

func (e TLUpdateMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateMessageID)
	x.Int(e.Id)
	x.Long(e.RandomId)
	return x.buf
}

func (e TLUpdateReadMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateReadMessages)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	return x.buf
}

func (e TLUpdateDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateDeleteMessages)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	return x.buf
}

func (e TLUpdateUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateUserTyping)
	x.Int(e.UserId)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TLUpdateChatUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateChatUserTyping)
	x.Int(e.ChatId)
	x.Int(e.UserId)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TLUpdateChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateChatParticipants)
	x.Bytes(e.Participants.encode())
	return x.buf
}

func (e TLUpdateUserStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateUserStatus)
	x.Int(e.UserId)
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e TLUpdateUserName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateUserName)
	x.Int(e.UserId)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Username)
	return x.buf
}

func (e TLUpdateUserPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateUserPhoto)
	x.Int(e.UserId)
	x.Int(e.Date)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Previous.encode())
	return x.buf
}

func (e TLUpdateContactRegistered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateContactRegistered)
	x.Int(e.UserId)
	x.Int(e.Date)
	return x.buf
}

func (e TLUpdateContactLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateContactLink)
	x.Int(e.UserId)
	x.Bytes(e.MyLink.encode())
	x.Bytes(e.ForeignLink.encode())
	return x.buf
}

func (e TLUpdateNewAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateNewAuthorization)
	x.Long(e.AuthKeyId)
	x.Int(e.Date)
	x.String(e.Device)
	x.String(e.Location)
	return x.buf
}

func (e TLUpdatesState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdatesState)
	x.Int(e.Pts)
	x.Int(e.Qts)
	x.Int(e.Date)
	x.Int(e.Seq)
	x.Int(e.UnreadCount)
	return x.buf
}

func (e TLUpdatesDifferenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdatesDifferenceEmpty)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

func (e TLUpdatesDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdatesDifference)
	x.Vector(e.NewMessages)
	x.Vector(e.NewEncryptedMessages)
	x.Vector(e.OtherUpdates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.encode())
	return x.buf
}

func (e TLUpdatesDifferenceSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdatesDifferenceSlice)
	x.Vector(e.NewMessages)
	x.Vector(e.NewEncryptedMessages)
	x.Vector(e.OtherUpdates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.IntermediateState.encode())
	return x.buf
}

func (e TLUpdatesTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdatesTooLong)
	return x.buf
}

func (e TLUpdateShortMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateShortMessage)
	x.Int(e.Id)
	x.Int(e.FromId)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

func (e TLUpdateShortChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateShortChatMessage)
	x.Int(e.Id)
	x.Int(e.FromId)
	x.Int(e.ChatId)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

func (e TLUpdateShort) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateShort)
	x.Bytes(e.Update.encode())
	x.Int(e.Date)
	return x.buf
}

func (e TLUpdatesCombined) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdatesCombined)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.SeqStart)
	x.Int(e.Seq)
	return x.buf
}

func (e TLUpdates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdates)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

func (e TLPhotosPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotosPhotos)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.buf
}

func (e TLPhotosPhotosSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotosPhotosSlice)
	x.Int(e.Count)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.buf
}

func (e TLPhotosPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotosPhoto)
	x.Bytes(e.Photo.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e TLUploadFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUploadFile)
	x.Bytes(e._type.encode())
	x.Int(e.Mtime)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TLDcOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDcOption)
	x.Int(e.Id)
	x.String(e.Hostname)
	x.String(e.IpAddress)
	x.Int(e.Port)
	return x.buf
}

func (e TLConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcConfig)
	x.Int(e.Date)
	x.Int(e.Expires)
	x.Bytes(e.TestMode.encode())
	x.Int(e.ThisDc)
	x.Vector(e.DcOptions)
	x.Int(e.ChatBigSize)
	x.Int(e.ChatSizeMax)
	x.Int(e.BroadcastSizeMax)
	x.Vector(e.DisabledFeatures)
	return x.buf
}

func (e TLNearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcNearestDc)
	x.String(e.Country)
	x.Int(e.ThisDc)
	x.Int(e.NearestDc)
	return x.buf
}

func (e TLHelpAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcHelpAppUpdate)
	x.Int(e.Id)
	x.Bytes(e.Critical.encode())
	x.String(e.Url)
	x.String(e.Text)
	return x.buf
}

func (e TLHelpNoAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcHelpNoAppUpdate)
	return x.buf
}

func (e TLHelpInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcHelpInviteText)
	x.String(e.Message)
	return x.buf
}

func (e TLMessagesStatedMessagesLinks) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesStatedMessagesLinks)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Vector(e.Links)
	x.Int(e.Pts)
	x.Int(e.Seq)
	return x.buf
}

func (e TLMessagesStatedMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesStatedMessageLink)
	x.Bytes(e.Message.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Vector(e.Links)
	x.Int(e.Pts)
	x.Int(e.Seq)
	return x.buf
}

func (e TLMessagesSentMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSentMessageLink)
	x.Int(e.Id)
	x.Int(e.Date)
	x.Int(e.Pts)
	x.Int(e.Seq)
	x.Vector(e.Links)
	return x.buf
}

func (e TLInputGeoChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputGeoChat)
	x.Int(e.ChatId)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLInputNotifyGeoChatPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputNotifyGeoChatPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TLGeoChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeoChat)
	x.Int(e.Id)
	x.Long(e.AccessHash)
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Venue)
	x.Bytes(e.Geo.encode())
	x.Bytes(e.Photo.encode())
	x.Int(e.ParticipantsCount)
	x.Int(e.Date)
	x.Bytes(e.CheckedIn.encode())
	x.Int(e.Version)
	return x.buf
}

func (e TLGeoChatMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeoChatMessageEmpty)
	x.Int(e.ChatId)
	x.Int(e.Id)
	return x.buf
}

func (e TLGeoChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeoChatMessage)
	x.Int(e.ChatId)
	x.Int(e.Id)
	x.Int(e.FromId)
	x.Int(e.Date)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	return x.buf
}

func (e TLGeoChatMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeoChatMessageService)
	x.Int(e.ChatId)
	x.Int(e.Id)
	x.Int(e.FromId)
	x.Int(e.Date)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TLGeochatsStatedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsStatedMessage)
	x.Bytes(e.Message.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Int(e.Seq)
	return x.buf
}

func (e TLGeochatsLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsLocated)
	x.Vector(e.Results)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TLGeochatsMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsMessages)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TLGeochatsMessagesSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsMessagesSlice)
	x.Int(e.Count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TLMessageActionGeoChatCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageActionGeoChatCreate)
	x.String(e.Title)
	x.String(e.Address)
	return x.buf
}

func (e TLMessageActionGeoChatCheckin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageActionGeoChatCheckin)
	return x.buf
}

func (e TLUpdateNewGeoChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateNewGeoChatMessage)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e TLWallPaperSolid) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcWallPaperSolid)
	x.Int(e.Id)
	x.String(e.Title)
	x.Int(e.BgColor)
	x.Int(e.Color)
	return x.buf
}

func (e TLUpdateNewEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateNewEncryptedMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Qts)
	return x.buf
}

func (e TLUpdateEncryptedChatTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateEncryptedChatTyping)
	x.Int(e.ChatId)
	return x.buf
}

func (e TLUpdateEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateEncryption)
	x.Bytes(e.Chat.encode())
	x.Int(e.Date)
	return x.buf
}

func (e TLUpdateEncryptedMessagesRead) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateEncryptedMessagesRead)
	x.Int(e.ChatId)
	x.Int(e.MaxDate)
	x.Int(e.Date)
	return x.buf
}

func (e TLEncryptedChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcEncryptedChatEmpty)
	x.Int(e.Id)
	return x.buf
}

func (e TLEncryptedChatWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcEncryptedChatWaiting)
	x.Int(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminId)
	x.Int(e.ParticipantId)
	return x.buf
}

func (e TLEncryptedChatRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcEncryptedChatRequested)
	x.Int(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminId)
	x.Int(e.ParticipantId)
	x.StringBytes(e.GA)
	return x.buf
}

func (e TLEncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcEncryptedChat)
	x.Int(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminId)
	x.Int(e.ParticipantId)
	x.StringBytes(e.GAOrB)
	x.Long(e.KeyFingerprint)
	return x.buf
}

func (e TLEncryptedChatDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcEncryptedChatDiscarded)
	x.Int(e.Id)
	return x.buf
}

func (e TLInputEncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputEncryptedChat)
	x.Int(e.ChatId)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLEncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcEncryptedFileEmpty)
	return x.buf
}

func (e TLEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcEncryptedFile)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Size)
	x.Int(e.DcId)
	x.Int(e.KeyFingerprint)
	return x.buf
}

func (e TLInputEncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputEncryptedFileEmpty)
	return x.buf
}

func (e TLInputEncryptedFileUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputEncryptedFileUploaded)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Md5Checksum)
	x.Int(e.KeyFingerprint)
	return x.buf
}

func (e TLInputEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputEncryptedFile)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLInputEncryptedFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputEncryptedFileLocation)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcEncryptedMessage)
	x.Long(e.RandomId)
	x.Int(e.ChatId)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TLEncryptedMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcEncryptedMessageService)
	x.Long(e.RandomId)
	x.Int(e.ChatId)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TLMessagesDhConfigNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesDhConfigNotModified)
	x.StringBytes(e.Random)
	return x.buf
}

func (e TLMessagesDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesDhConfig)
	x.Int(e.G)
	x.StringBytes(e.P)
	x.Int(e.Version)
	x.StringBytes(e.Random)
	return x.buf
}

func (e TLMessagesSentEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSentEncryptedMessage)
	x.Int(e.Date)
	return x.buf
}

func (e TLMessagesSentEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSentEncryptedFile)
	x.Int(e.Date)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TLInputFileBig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputFileBig)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Name)
	return x.buf
}

func (e TLInputEncryptedFileBigUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputEncryptedFileBigUploaded)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.Int(e.KeyFingerprint)
	return x.buf
}

func (e TLUpdateChatParticipantAdd) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateChatParticipantAdd)
	x.Int(e.ChatId)
	x.Int(e.UserId)
	x.Int(e.InviterId)
	x.Int(e.Version)
	return x.buf
}

func (e TLUpdateChatParticipantDelete) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateChatParticipantDelete)
	x.Int(e.ChatId)
	x.Int(e.UserId)
	x.Int(e.Version)
	return x.buf
}

func (e TLUpdateDcOptions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateDcOptions)
	x.Vector(e.DcOptions)
	return x.buf
}

func (e TLInputMediaUploadedAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaUploadedAudio)
	x.Bytes(e.File.encode())
	x.Int(e.Duration)
	x.String(e.MimeType)
	return x.buf
}

func (e TLInputMediaAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaAudio)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TLInputMediaUploadedDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaUploadedDocument)
	x.Bytes(e.File.encode())
	x.String(e.MimeType)
	x.Vector(e.Attributes)
	return x.buf
}

func (e TLInputMediaUploadedThumbDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaUploadedThumbDocument)
	x.Bytes(e.File.encode())
	x.Bytes(e.Thumb.encode())
	x.String(e.MimeType)
	x.Vector(e.Attributes)
	return x.buf
}

func (e TLInputMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputMediaDocument)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TLMessageMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageMediaDocument)
	x.Bytes(e.Document.encode())
	return x.buf
}

func (e TLMessageMediaAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessageMediaAudio)
	x.Bytes(e.Audio.encode())
	return x.buf
}

func (e TLInputAudioEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputAudioEmpty)
	return x.buf
}

func (e TLInputAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputAudio)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLInputDocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputDocumentEmpty)
	return x.buf
}

func (e TLInputDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputDocument)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLInputAudioFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputAudioFileLocation)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLInputDocumentFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputDocumentFileLocation)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TLAudioEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAudioEmpty)
	x.Long(e.Id)
	return x.buf
}

func (e TLAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAudio)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.UserId)
	x.Int(e.Date)
	x.Int(e.Duration)
	x.String(e.MimeType)
	x.Int(e.Size)
	x.Int(e.DcId)
	return x.buf
}

func (e TLDocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDocumentEmpty)
	x.Long(e.Id)
	return x.buf
}

func (e TLDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDocument)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.String(e.MimeType)
	x.Int(e.Size)
	x.Bytes(e.Thumb.encode())
	x.Int(e.DcId)
	x.Vector(e.Attributes)
	return x.buf
}

func (e TLHelpSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcHelpSupport)
	x.String(e.PhoneNumber)
	x.Bytes(e.User.encode())
	return x.buf
}

func (e TLNotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcNotifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TLNotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcNotifyUsers)
	return x.buf
}

func (e TLNotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcNotifyChats)
	return x.buf
}

func (e TLNotifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcNotifyAll)
	return x.buf
}

func (e TLUpdateUserBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateUserBlocked)
	x.Int(e.UserId)
	x.Bytes(e.Blocked.encode())
	return x.buf
}

func (e TLUpdateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.NotifySettings.encode())
	return x.buf
}

func (e TLAuthSentAppCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthSentAppCode)
	x.Bytes(e.PhoneRegistered.encode())
	x.String(e.PhoneCodeHash)
	x.Int(e.SendCallTimeout)
	x.Bytes(e.IsPassword.encode())
	return x.buf
}

func (e TLSendMessageTypingAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcSendMessageTypingAction)
	return x.buf
}

func (e TLSendMessageCancelAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcSendMessageCancelAction)
	return x.buf
}

func (e TLSendMessageRecordVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcSendMessageRecordVideoAction)
	return x.buf
}

func (e TLSendMessageUploadVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcSendMessageUploadVideoAction)
	return x.buf
}

func (e TLSendMessageRecordAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcSendMessageRecordAudioAction)
	return x.buf
}

func (e TLSendMessageUploadAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcSendMessageUploadAudioAction)
	return x.buf
}

func (e TLSendMessageUploadPhotoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcSendMessageUploadPhotoAction)
	return x.buf
}

func (e TLSendMessageUploadDocumentAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcSendMessageUploadDocumentAction)
	return x.buf
}

func (e TLSendMessageGeoLocationAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcSendMessageGeoLocationAction)
	return x.buf
}

func (e TLSendMessageChooseContactAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcSendMessageChooseContactAction)
	return x.buf
}

func (e TLContactFound) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactFound)
	x.Int(e.UserId)
	return x.buf
}

func (e TLContactsFound) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsFound)
	x.Vector(e.Results)
	x.Vector(e.Users)
	return x.buf
}

func (e TLUpdateServiceNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateServiceNotification)
	x.String(e._type)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	x.Bytes(e.Popup.encode())
	return x.buf
}

func (e TLUserStatusRecently) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserStatusRecently)
	return x.buf
}

func (e TLUserStatusLastWeek) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserStatusLastWeek)
	return x.buf
}

func (e TLUserStatusLastMonth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUserStatusLastMonth)
	return x.buf
}

func (e TLUpdatePrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdatePrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(e.Rules)
	return x.buf
}

func (e TLInputPrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPrivacyKeyStatusTimestamp)
	return x.buf
}

func (e TLPrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPrivacyKeyStatusTimestamp)
	return x.buf
}

func (e TLInputPrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPrivacyValueAllowContacts)
	return x.buf
}

func (e TLInputPrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPrivacyValueAllowAll)
	return x.buf
}

func (e TLInputPrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPrivacyValueAllowUsers)
	x.Vector(e.Users)
	return x.buf
}

func (e TLInputPrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPrivacyValueDisallowContacts)
	return x.buf
}

func (e TLInputPrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPrivacyValueDisallowAll)
	return x.buf
}

func (e TLInputPrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInputPrivacyValueDisallowUsers)
	x.Vector(e.Users)
	return x.buf
}

func (e TLPrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPrivacyValueAllowContacts)
	return x.buf
}

func (e TLPrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPrivacyValueAllowAll)
	return x.buf
}

func (e TLPrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPrivacyValueAllowUsers)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TLPrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPrivacyValueDisallowContacts)
	return x.buf
}

func (e TLPrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPrivacyValueDisallowAll)
	return x.buf
}

func (e TLPrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPrivacyValueDisallowUsers)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TLAccountPrivacyRules) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountPrivacyRules)
	x.Vector(e.Rules)
	x.Vector(e.Users)
	return x.buf
}

func (e TLAccountDaysTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountDaysTTL)
	x.Int(e.Days)
	return x.buf
}

func (e TLAccountSentChangePhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountSentChangePhoneCode)
	x.String(e.PhoneCodeHash)
	x.Int(e.SendCallTimeout)
	return x.buf
}

func (e TLUpdateUserPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdateUserPhone)
	x.Int(e.UserId)
	x.String(e.Phone)
	return x.buf
}

func (e TLDocumentAttributeImageSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDocumentAttributeImageSize)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TLDocumentAttributeAnimated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDocumentAttributeAnimated)
	return x.buf
}

func (e TLDocumentAttributeSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDocumentAttributeSticker)
	return x.buf
}

func (e TLDocumentAttributeVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDocumentAttributeVideo)
	x.Int(e.Duration)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TLDocumentAttributeAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDocumentAttributeAudio)
	x.Int(e.Duration)
	return x.buf
}

func (e TLDocumentAttributeFilename) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDocumentAttributeFilename)
	x.String(e.FileName)
	return x.buf
}

func (e TLMessagesStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesStickersNotModified)
	return x.buf
}

func (e TLMessagesStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesStickers)
	x.String(e.Hash)
	x.Vector(e.Stickers)
	return x.buf
}

func (e TLStickerPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcStickerPack)
	x.String(e.Emoticon)
	x.VectorLong(e.Documents)
	return x.buf
}

func (e TLMessagesAllStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesAllStickersNotModified)
	return x.buf
}

func (e TLMessagesAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesAllStickers)
	x.String(e.Hash)
	x.Vector(e.Packs)
	x.Vector(e.Documents)
	return x.buf
}

func (e TLDisabledFeature) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcDisabledFeature)
	x.String(e.Feature)
	x.String(e.Description)
	return x.buf
}

func (e TLInvokeAfterMsg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInvokeAfterMsg)
	x.Long(e.MsgId)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TLInvokeAfterMsgs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInvokeAfterMsgs)
	x.VectorLong(e.MsgIds)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TLAuthCheckPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthCheckPhone)
	x.String(e.PhoneNumber)
	return x.buf
}

func (e TLAuthSendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthSendCode)
	x.String(e.PhoneNumber)
	x.Int(e.SmsType)
	x.Int(e.ApiId)
	x.String(e.ApiHash)
	x.String(e.LangCode)
	return x.buf
}

func (e TLAuthSendCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthSendCall)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	return x.buf
}

func (e TLAuthSignUp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthSignUp)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}

func (e TLAuthSignIn) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthSignIn)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}

func (e TLAuthLogOut) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthLogOut)
	return x.buf
}

func (e TLAuthResetAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthResetAuthorizations)
	return x.buf
}

func (e TLAuthSendInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthSendInvites)
	x.VectorString(e.PhoneNumbers)
	x.String(e.Message)
	return x.buf
}

func (e TLAuthExportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthExportAuthorization)
	x.Int(e.DcId)
	return x.buf
}

func (e TLAuthImportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthImportAuthorization)
	x.Int(e.Id)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TLAuthBindTempAuthKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthBindTempAuthKey)
	x.Long(e.PermAuthKeyId)
	x.Long(e.Nonce)
	x.Int(e.ExpiresAt)
	x.StringBytes(e.EncryptedMessage)
	return x.buf
}

func (e TLAccountRegisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountRegisterDevice)
	x.Int(e.TokenType)
	x.String(e.Token)
	x.String(e.DeviceModel)
	x.String(e.SystemVersion)
	x.String(e.AppVersion)
	x.Bytes(e.AppSandbox.encode())
	x.String(e.LangCode)
	return x.buf
}

func (e TLAccountUnregisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountUnregisterDevice)
	x.Int(e.TokenType)
	x.String(e.Token)
	return x.buf
}

func (e TLAccountUpdateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountUpdateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TLAccountGetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountGetNotifySettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TLAccountResetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountResetNotifySettings)
	return x.buf
}

func (e TLAccountUpdateProfile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountUpdateProfile)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}

func (e TLAccountUpdateStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountUpdateStatus)
	x.Bytes(e.Offline.encode())
	return x.buf
}

func (e TLAccountGetWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountGetWallPapers)
	return x.buf
}

func (e TLUsersGetUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUsersGetUsers)
	x.Vector(e.Id)
	return x.buf
}

func (e TLUsersGetFullUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUsersGetFullUser)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TLContactsGetStatuses) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsGetStatuses)
	return x.buf
}

func (e TLContactsGetContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsGetContacts)
	x.String(e.Hash)
	return x.buf
}

func (e TLContactsImportContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsImportContacts)
	x.Vector(e.Contacts)
	x.Bytes(e.Replace.encode())
	return x.buf
}

func (e TLContactsGetSuggested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsGetSuggested)
	x.Int(e.Limit)
	return x.buf
}

func (e TLContactsDeleteContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsDeleteContact)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TLContactsDeleteContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsDeleteContacts)
	x.Vector(e.Id)
	return x.buf
}

func (e TLContactsBlock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsBlock)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TLContactsUnblock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsUnblock)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TLContactsGetBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsGetBlocked)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TLContactsExportCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsExportCard)
	return x.buf
}

func (e TLContactsImportCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsImportCard)
	x.VectorInt(e.ExportCard)
	return x.buf
}

func (e TLMessagesGetMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesGetMessages)
	x.VectorInt(e.Id)
	return x.buf
}

func (e TLMessagesGetDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesGetDialogs)
	x.Int(e.Offset)
	x.Int(e.MaxId)
	x.Int(e.Limit)
	return x.buf
}

func (e TLMessagesGetHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesGetHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Offset)
	x.Int(e.MaxId)
	x.Int(e.Limit)
	return x.buf
}

func (e TLMessagesSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSearch)
	x.Bytes(e.Peer.encode())
	x.String(e.Q)
	x.Bytes(e.Filter.encode())
	x.Int(e.MinDate)
	x.Int(e.MaxDate)
	x.Int(e.Offset)
	x.Int(e.MaxId)
	x.Int(e.Limit)
	return x.buf
}

func (e TLMessagesReadHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesReadHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxId)
	x.Int(e.Offset)
	x.Bytes(e.ReadContents.encode())
	return x.buf
}

func (e TLMessagesDeleteHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesDeleteHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Offset)
	return x.buf
}

func (e TLMessagesDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesDeleteMessages)
	x.VectorInt(e.Id)
	return x.buf
}

func (e TLMessagesReceivedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesReceivedMessages)
	x.Int(e.MaxId)
	return x.buf
}

func (e TLMessagesSetTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSetTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TLMessagesSendMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSendMessage)
	x.Bytes(e.Peer.encode())
	x.String(e.Message)
	x.Long(e.RandomId)
	return x.buf
}

func (e TLMessagesSendMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSendMedia)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Media.encode())
	x.Long(e.RandomId)
	return x.buf
}

func (e TLMessagesForwardMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesForwardMessages)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.Id)
	return x.buf
}

func (e TLMessagesGetChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesGetChats)
	x.VectorInt(e.Id)
	return x.buf
}

func (e TLMessagesGetFullChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesGetFullChat)
	x.Int(e.ChatId)
	return x.buf
}

func (e TLMessagesEditChatTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesEditChatTitle)
	x.Int(e.ChatId)
	x.String(e.Title)
	return x.buf
}

func (e TLMessagesEditChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesEditChatPhoto)
	x.Int(e.ChatId)
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e TLMessagesAddChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesAddChatUser)
	x.Int(e.ChatId)
	x.Bytes(e.UserId.encode())
	x.Int(e.FwdLimit)
	return x.buf
}

func (e TLMessagesDeleteChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesDeleteChatUser)
	x.Int(e.ChatId)
	x.Bytes(e.UserId.encode())
	return x.buf
}

func (e TLMessagesCreateChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesCreateChat)
	x.Vector(e.Users)
	x.String(e.Title)
	return x.buf
}

func (e TLUpdatesGetState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdatesGetState)
	return x.buf
}

func (e TLUpdatesGetDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUpdatesGetDifference)
	x.Int(e.Pts)
	x.Int(e.Date)
	x.Int(e.Qts)
	return x.buf
}

func (e TLPhotosUpdateProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotosUpdateProfilePhoto)
	x.Bytes(e.Id.encode())
	x.Bytes(e.Crop.encode())
	return x.buf
}

func (e TLPhotosUploadProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotosUploadProfilePhoto)
	x.Bytes(e.File.encode())
	x.String(e.Caption)
	x.Bytes(e.GeoPoint.encode())
	x.Bytes(e.Crop.encode())
	return x.buf
}

func (e TLPhotosDeletePhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotosDeletePhotos)
	x.Vector(e.Id)
	return x.buf
}

func (e TLUploadSaveFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUploadSaveFilePart)
	x.Long(e.FileId)
	x.Int(e.FilePart)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TLUploadGetFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUploadGetFile)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TLHelpGetConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcHelpGetConfig)
	return x.buf
}

func (e TLHelpGetNearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcHelpGetNearestDc)
	return x.buf
}

func (e TLHelpGetAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcHelpGetAppUpdate)
	x.String(e.DeviceModel)
	x.String(e.SystemVersion)
	x.String(e.AppVersion)
	x.String(e.LangCode)
	return x.buf
}

func (e TLHelpSaveAppLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcHelpSaveAppLog)
	x.Vector(e.Events)
	return x.buf
}

func (e TLHelpGetInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcHelpGetInviteText)
	x.String(e.LangCode)
	return x.buf
}

func (e TLPhotosGetUserPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcPhotosGetUserPhotos)
	x.Bytes(e.UserId.encode())
	x.Int(e.Offset)
	x.Int(e.MaxId)
	x.Int(e.Limit)
	return x.buf
}

func (e TLMessagesForwardMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesForwardMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.Long(e.RandomId)
	return x.buf
}

func (e TLMessagesSendBroadcast) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSendBroadcast)
	x.Vector(e.Contacts)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	return x.buf
}

func (e TLGeochatsGetLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsGetLocated)
	x.Bytes(e.GeoPoint.encode())
	x.Int(e.Radius)
	x.Int(e.Limit)
	return x.buf
}

func (e TLGeochatsGetRecents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsGetRecents)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TLGeochatsCheckin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsCheckin)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TLGeochatsGetFullChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsGetFullChat)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TLGeochatsEditChatTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsEditChatTitle)
	x.Bytes(e.Peer.encode())
	x.String(e.Title)
	x.String(e.Address)
	return x.buf
}

func (e TLGeochatsEditChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsEditChatPhoto)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e TLGeochatsSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsSearch)
	x.Bytes(e.Peer.encode())
	x.String(e.Q)
	x.Bytes(e.Filter.encode())
	x.Int(e.MinDate)
	x.Int(e.MaxDate)
	x.Int(e.Offset)
	x.Int(e.MaxId)
	x.Int(e.Limit)
	return x.buf
}

func (e TLGeochatsGetHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsGetHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Offset)
	x.Int(e.MaxId)
	x.Int(e.Limit)
	return x.buf
}

func (e TLGeochatsSetTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsSetTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Typing.encode())
	return x.buf
}

func (e TLGeochatsSendMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsSendMessage)
	x.Bytes(e.Peer.encode())
	x.String(e.Message)
	x.Long(e.RandomId)
	return x.buf
}

func (e TLGeochatsSendMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsSendMedia)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Media.encode())
	x.Long(e.RandomId)
	return x.buf
}

func (e TLGeochatsCreateGeoChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcGeochatsCreateGeoChat)
	x.String(e.Title)
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Address)
	x.String(e.Venue)
	return x.buf
}

func (e TLMessagesGetDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesGetDhConfig)
	x.Int(e.Version)
	x.Int(e.RandomLength)
	return x.buf
}

func (e TLMessagesRequestEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesRequestEncryption)
	x.Bytes(e.UserId.encode())
	x.Int(e.RandomId)
	x.StringBytes(e.GA)
	return x.buf
}

func (e TLMessagesAcceptEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesAcceptEncryption)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.GB)
	x.Long(e.KeyFingerprint)
	return x.buf
}

func (e TLMessagesDiscardEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesDiscardEncryption)
	x.Int(e.ChatId)
	return x.buf
}

func (e TLMessagesSetEncryptedTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSetEncryptedTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Typing.encode())
	return x.buf
}

func (e TLMessagesReadEncryptedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesReadEncryptedHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxDate)
	return x.buf
}

func (e TLMessagesSendEncrypted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSendEncrypted)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomId)
	x.StringBytes(e.Data)
	return x.buf
}

func (e TLMessagesSendEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSendEncryptedFile)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomId)
	x.StringBytes(e.Data)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TLMessagesSendEncryptedService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesSendEncryptedService)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomId)
	x.StringBytes(e.Data)
	return x.buf
}

func (e TLMessagesReceivedQueue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesReceivedQueue)
	x.Int(e.MaxQts)
	return x.buf
}

func (e TLUploadSaveBigFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcUploadSaveBigFilePart)
	x.Long(e.FileId)
	x.Int(e.FilePart)
	x.Int(e.FileTotalParts)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TLInitConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInitConnection)
	x.Int(e.ApiId)
	x.String(e.DeviceModel)
	x.String(e.SystemVersion)
	x.String(e.AppVersion)
	x.String(e.LangCode)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TLHelpGetSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcHelpGetSupport)
	return x.buf
}

func (e TLAuthSendSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAuthSendSms)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	return x.buf
}

func (e TLMessagesReadMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesReadMessageContents)
	x.VectorInt(e.Id)
	return x.buf
}

func (e TLAccountCheckUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountCheckUsername)
	x.String(e.Username)
	return x.buf
}

func (e TLAccountUpdateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountUpdateUsername)
	x.String(e.Username)
	return x.buf
}

func (e TLContactsSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsSearch)
	x.String(e.Q)
	x.Int(e.Limit)
	return x.buf
}

func (e TLAccountGetPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountGetPrivacy)
	x.Bytes(e.Key.encode())
	return x.buf
}

func (e TLAccountSetPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountSetPrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(e.Rules)
	return x.buf
}

func (e TLAccountDeleteAccount) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountDeleteAccount)
	x.String(e.Reason)
	return x.buf
}

func (e TLAccountGetAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountGetAccountTTL)
	return x.buf
}

func (e TLAccountSetAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountSetAccountTTL)
	x.Bytes(e.Ttl.encode())
	return x.buf
}

func (e TLInvokeWithLayer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcInvokeWithLayer)
	x.Int(e.Layer)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TLContactsResolveUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcContactsResolveUsername)
	x.String(e.Username)
	return x.buf
}

func (e TLAccountSendChangePhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountSendChangePhoneCode)
	x.String(e.PhoneNumber)
	return x.buf
}

func (e TLAccountChangePhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountChangePhone)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}

func (e TLMessagesGetStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesGetStickers)
	x.String(e.Emoticon)
	x.String(e.Hash)
	return x.buf
}

func (e TLMessagesGetAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcMessagesGetAllStickers)
	x.String(e.Hash)
	return x.buf
}

func (e TLAccountUpdateDeviceLocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcAccountUpdateDeviceLocked)
	x.Int(e.Period)
	return x.buf
}

func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {
	case CrcBoolFalse:
		r = TLBoolFalse{}

	case CrcBoolTrue:
		r = TLBoolTrue{}

	case CrcVector:
		r = TLVector{}

	case CrcError:
		r = TLError{
			m.Int(),
			m.String(),
		}

	case CrcNull:
		r = TLNull{}

	case CrcInputPeerEmpty:
		r = TLInputPeerEmpty{}

	case CrcInputPeerSelf:
		r = TLInputPeerSelf{}

	case CrcInputPeerContact:
		r = TLInputPeerContact{
			m.Int(),
		}

	case CrcInputPeerForeign:
		r = TLInputPeerForeign{
			m.Int(),
			m.Long(),
		}

	case CrcInputPeerChat:
		r = TLInputPeerChat{
			m.Int(),
		}

	case CrcInputUserEmpty:
		r = TLInputUserEmpty{}

	case CrcInputUserSelf:
		r = TLInputUserSelf{}

	case CrcInputUserContact:
		r = TLInputUserContact{
			m.Int(),
		}

	case CrcInputUserForeign:
		r = TLInputUserForeign{
			m.Int(),
			m.Long(),
		}

	case CrcInputPhoneContact:
		r = TLInputPhoneContact{
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CrcInputFile:
		r = TLInputFile{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case CrcInputMediaEmpty:
		r = TLInputMediaEmpty{}

	case CrcInputMediaUploadedPhoto:
		r = TLInputMediaUploadedPhoto{
			m.Object(),
		}

	case CrcInputMediaPhoto:
		r = TLInputMediaPhoto{
			m.Object(),
		}

	case CrcInputMediaGeoPoint:
		r = TLInputMediaGeoPoint{
			m.Object(),
		}

	case CrcInputMediaContact:
		r = TLInputMediaContact{
			m.String(),
			m.String(),
			m.String(),
		}

	case CrcInputMediaUploadedVideo:
		r = TLInputMediaUploadedVideo{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
		}

	case CrcInputMediaUploadedThumbVideo:
		r = TLInputMediaUploadedThumbVideo{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
		}

	case CrcInputMediaVideo:
		r = TLInputMediaVideo{
			m.Object(),
		}

	case CrcInputChatPhotoEmpty:
		r = TLInputChatPhotoEmpty{}

	case CrcInputChatUploadedPhoto:
		r = TLInputChatUploadedPhoto{
			m.Object(),
			m.Object(),
		}

	case CrcInputChatPhoto:
		r = TLInputChatPhoto{
			m.Object(),
			m.Object(),
		}

	case CrcInputGeoPointEmpty:
		r = TLInputGeoPointEmpty{}

	case CrcInputGeoPoint:
		r = TLInputGeoPoint{
			m.Double(),
			m.Double(),
		}

	case CrcInputPhotoEmpty:
		r = TLInputPhotoEmpty{}

	case CrcInputPhoto:
		r = TLInputPhoto{
			m.Long(),
			m.Long(),
		}

	case CrcInputVideoEmpty:
		r = TLInputVideoEmpty{}

	case CrcInputVideo:
		r = TLInputVideo{
			m.Long(),
			m.Long(),
		}

	case CrcInputFileLocation:
		r = TLInputFileLocation{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case CrcInputVideoFileLocation:
		r = TLInputVideoFileLocation{
			m.Long(),
			m.Long(),
		}

	case CrcInputPhotoCropAuto:
		r = TLInputPhotoCropAuto{}

	case CrcInputPhotoCrop:
		r = TLInputPhotoCrop{
			m.Double(),
			m.Double(),
			m.Double(),
		}

	case CrcInputAppEvent:
		r = TLInputAppEvent{
			m.Double(),
			m.String(),
			m.Long(),
			m.String(),
		}

	case CrcPeerUser:
		r = TLPeerUser{
			m.Int(),
		}

	case CrcPeerChat:
		r = TLPeerChat{
			m.Int(),
		}

	case CrcStorageFileUnknown:
		r = TLStorageFileUnknown{}

	case CrcStorageFileJpeg:
		r = TLStorageFileJpeg{}

	case CrcStorageFileGif:
		r = TLStorageFileGif{}

	case CrcStorageFilePng:
		r = TLStorageFilePng{}

	case CrcStorageFilePdf:
		r = TLStorageFilePdf{}

	case CrcStorageFileMp3:
		r = TLStorageFileMp3{}

	case CrcStorageFileMov:
		r = TLStorageFileMov{}

	case CrcStorageFilePartial:
		r = TLStorageFilePartial{}

	case CrcStorageFileMp4:
		r = TLStorageFileMp4{}

	case CrcStorageFileWebp:
		r = TLStorageFileWebp{}

	case CrcFileLocationUnavailable:
		r = TLFileLocationUnavailable{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case CrcFileLocation:
		r = TLFileLocation{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case CrcUserEmpty:
		r = TLUserEmpty{
			m.Int(),
		}

	case CrcUserSelf:
		r = TLUserSelf{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case CrcUserContact:
		r = TLUserContact{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Long(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case CrcUserRequest:
		r = TLUserRequest{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Long(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case CrcUserForeign:
		r = TLUserForeign{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case CrcUserDeleted:
		r = TLUserDeleted{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CrcUserProfilePhotoEmpty:
		r = TLUserProfilePhotoEmpty{}

	case CrcUserProfilePhoto:
		r = TLUserProfilePhoto{
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case CrcUserStatusEmpty:
		r = TLUserStatusEmpty{}

	case CrcUserStatusOnline:
		r = TLUserStatusOnline{
			m.Int(),
		}

	case CrcUserStatusOffline:
		r = TLUserStatusOffline{
			m.Int(),
		}

	case CrcChatEmpty:
		r = TLChatEmpty{
			m.Int(),
		}

	case CrcChat:
		r = TLChat{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case CrcChatForbidden:
		r = TLChatForbidden{
			m.Int(),
			m.String(),
			m.Int(),
		}

	case CrcChatFull:
		r = TLChatFull{
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case CrcChatParticipant:
		r = TLChatParticipant{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcChatParticipantsForbidden:
		r = TLChatParticipantsForbidden{
			m.Int(),
		}

	case CrcChatParticipants:
		r = TLChatParticipants{
			m.Int(),
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case CrcChatPhotoEmpty:
		r = TLChatPhotoEmpty{}

	case CrcChatPhoto:
		r = TLChatPhoto{
			m.Object(),
			m.Object(),
		}

	case CrcMessageEmpty:
		r = TLMessageEmpty{
			m.Int(),
		}

	case CrcMessage:
		r = TLMessage{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
		}

	case CrcMessageForwarded:
		r = TLMessageForwarded{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
		}

	case CrcMessageService:
		r = TLMessageService{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case CrcMessageMediaEmpty:
		r = TLMessageMediaEmpty{}

	case CrcMessageMediaPhoto:
		r = TLMessageMediaPhoto{
			m.Object(),
		}

	case CrcMessageMediaVideo:
		r = TLMessageMediaVideo{
			m.Object(),
		}

	case CrcMessageMediaGeo:
		r = TLMessageMediaGeo{
			m.Object(),
		}

	case CrcMessageMediaContact:
		r = TLMessageMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case CrcMessageMediaUnsupported:
		r = TLMessageMediaUnsupported{
			m.StringBytes(),
		}

	case CrcMessageActionEmpty:
		r = TLMessageActionEmpty{}

	case CrcMessageActionChatCreate:
		r = TLMessageActionChatCreate{
			m.String(),
			m.VectorInt(),
		}

	case CrcMessageActionChatEditTitle:
		r = TLMessageActionChatEditTitle{
			m.String(),
		}

	case CrcMessageActionChatEditPhoto:
		r = TLMessageActionChatEditPhoto{
			m.Object(),
		}

	case CrcMessageActionChatDeletePhoto:
		r = TLMessageActionChatDeletePhoto{}

	case CrcMessageActionChatAddUser:
		r = TLMessageActionChatAddUser{
			m.Int(),
		}

	case CrcMessageActionChatDeleteUser:
		r = TLMessageActionChatDeleteUser{
			m.Int(),
		}

	case CrcDialog:
		r = TLDialog{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case CrcPhotoEmpty:
		r = TLPhotoEmpty{
			m.Long(),
		}

	case CrcPhoto:
		r = TLPhoto{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Object(),
			m.Vector(),
		}

	case CrcPhotoSizeEmpty:
		r = TLPhotoSizeEmpty{
			m.String(),
		}

	case CrcPhotoSize:
		r = TLPhotoSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcPhotoCachedSize:
		r = TLPhotoCachedSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case CrcVideoEmpty:
		r = TLVideoEmpty{
			m.Long(),
		}

	case CrcVideo:
		r = TLVideo{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcGeoPointEmpty:
		r = TLGeoPointEmpty{}

	case CrcGeoPoint:
		r = TLGeoPoint{
			m.Double(),
			m.Double(),
		}

	case CrcAuthCheckedPhone:
		r = TLAuthCheckedPhone{
			m.Object(),
			m.Object(),
		}

	case CrcAuthSentCode:
		r = TLAuthSentCode{
			m.Object(),
			m.String(),
			m.Int(),
			m.Object(),
		}

	case CrcAuthAuthorization:
		r = TLAuthAuthorization{
			m.Int(),
			m.Object(),
		}

	case CrcAuthExportedAuthorization:
		r = TLAuthExportedAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case CrcInputNotifyPeer:
		r = TLInputNotifyPeer{
			m.Object(),
		}

	case CrcInputNotifyUsers:
		r = TLInputNotifyUsers{}

	case CrcInputNotifyChats:
		r = TLInputNotifyChats{}

	case CrcInputNotifyAll:
		r = TLInputNotifyAll{}

	case CrcInputPeerNotifyEventsEmpty:
		r = TLInputPeerNotifyEventsEmpty{}

	case CrcInputPeerNotifyEventsAll:
		r = TLInputPeerNotifyEventsAll{}

	case CrcInputPeerNotifySettings:
		r = TLInputPeerNotifySettings{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
		}

	case CrcPeerNotifyEventsEmpty:
		r = TLPeerNotifyEventsEmpty{}

	case CrcPeerNotifyEventsAll:
		r = TLPeerNotifyEventsAll{}

	case CrcPeerNotifySettingsEmpty:
		r = TLPeerNotifySettingsEmpty{}

	case CrcPeerNotifySettings:
		r = TLPeerNotifySettings{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
		}

	case CrcWallPaper:
		r = TLWallPaper{
			m.Int(),
			m.String(),
			m.Vector(),
			m.Int(),
		}

	case CrcUserFull:
		r = TLUserFull{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case CrcContact:
		r = TLContact{
			m.Int(),
			m.Object(),
		}

	case CrcImportedContact:
		r = TLImportedContact{
			m.Int(),
			m.Long(),
		}

	case CrcContactBlocked:
		r = TLContactBlocked{
			m.Int(),
			m.Int(),
		}

	case CrcContactSuggested:
		r = TLContactSuggested{
			m.Int(),
			m.Int(),
		}

	case CrcContactStatus:
		r = TLContactStatus{
			m.Int(),
			m.Object(),
		}

	case CrcChatLocated:
		r = TLChatLocated{
			m.Int(),
			m.Int(),
		}

	case CrcContactsForeignLinkUnknown:
		r = TLContactsForeignLinkUnknown{}

	case CrcContactsForeignLinkRequested:
		r = TLContactsForeignLinkRequested{
			m.Object(),
		}

	case CrcContactsForeignLinkMutual:
		r = TLContactsForeignLinkMutual{}

	case CrcContactsMyLinkEmpty:
		r = TLContactsMyLinkEmpty{}

	case CrcContactsMyLinkRequested:
		r = TLContactsMyLinkRequested{
			m.Object(),
		}

	case CrcContactsMyLinkContact:
		r = TLContactsMyLinkContact{}

	case CrcContactsLink:
		r = TLContactsLink{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case CrcContactsContactsNotModified:
		r = TLContactsContactsNotModified{}

	case CrcContactsContacts:
		r = TLContactsContacts{
			m.Vector(),
			m.Vector(),
		}

	case CrcContactsImportedContacts:
		r = TLContactsImportedContacts{
			m.Vector(),
			m.VectorLong(),
			m.Vector(),
		}

	case CrcContactsBlocked:
		r = TLContactsBlocked{
			m.Vector(),
			m.Vector(),
		}

	case CrcContactsBlockedSlice:
		r = TLContactsBlockedSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case CrcContactsSuggested:
		r = TLContactsSuggested{
			m.Vector(),
			m.Vector(),
		}

	case CrcMessagesDialogs:
		r = TLMessagesDialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CrcMessagesDialogsSlice:
		r = TLMessagesDialogsSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CrcMessagesMessages:
		r = TLMessagesMessages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CrcMessagesMessagesSlice:
		r = TLMessagesMessagesSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CrcMessagesMessageEmpty:
		r = TLMessagesMessageEmpty{}

	case CrcMessagesStatedMessages:
		r = TLMessagesStatedMessages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case CrcMessagesStatedMessage:
		r = TLMessagesStatedMessage{
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case CrcMessagesSentMessage:
		r = TLMessagesSentMessage{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcMessagesChats:
		r = TLMessagesChats{
			m.Vector(),
			m.Vector(),
		}

	case CrcMessagesChatFull:
		r = TLMessagesChatFull{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case CrcMessagesAffectedHistory:
		r = TLMessagesAffectedHistory{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcInputMessagesFilterEmpty:
		r = TLInputMessagesFilterEmpty{}

	case CrcInputMessagesFilterPhotos:
		r = TLInputMessagesFilterPhotos{}

	case CrcInputMessagesFilterVideo:
		r = TLInputMessagesFilterVideo{}

	case CrcInputMessagesFilterPhotoVideo:
		r = TLInputMessagesFilterPhotoVideo{}

	case CrcInputMessagesFilterPhotoVideoDocuments:
		r = TLInputMessagesFilterPhotoVideoDocuments{}

	case CrcInputMessagesFilterDocument:
		r = TLInputMessagesFilterDocument{}

	case CrcInputMessagesFilterAudio:
		r = TLInputMessagesFilterAudio{}

	case CrcUpdateNewMessage:
		r = TLUpdateNewMessage{
			m.Object(),
			m.Int(),
		}

	case CrcUpdateMessageID:
		r = TLUpdateMessageID{
			m.Int(),
			m.Long(),
		}

	case CrcUpdateReadMessages:
		r = TLUpdateReadMessages{
			m.VectorInt(),
			m.Int(),
		}

	case CrcUpdateDeleteMessages:
		r = TLUpdateDeleteMessages{
			m.VectorInt(),
			m.Int(),
		}

	case CrcUpdateUserTyping:
		r = TLUpdateUserTyping{
			m.Int(),
			m.Object(),
		}

	case CrcUpdateChatUserTyping:
		r = TLUpdateChatUserTyping{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case CrcUpdateChatParticipants:
		r = TLUpdateChatParticipants{
			m.Object(),
		}

	case CrcUpdateUserStatus:
		r = TLUpdateUserStatus{
			m.Int(),
			m.Object(),
		}

	case CrcUpdateUserName:
		r = TLUpdateUserName{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CrcUpdateUserPhoto:
		r = TLUpdateUserPhoto{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case CrcUpdateContactRegistered:
		r = TLUpdateContactRegistered{
			m.Int(),
			m.Int(),
		}

	case CrcUpdateContactLink:
		r = TLUpdateContactLink{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case CrcUpdateNewAuthorization:
		r = TLUpdateNewAuthorization{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case CrcUpdatesState:
		r = TLUpdatesState{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcUpdatesDifferenceEmpty:
		r = TLUpdatesDifferenceEmpty{
			m.Int(),
			m.Int(),
		}

	case CrcUpdatesDifference:
		r = TLUpdatesDifference{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case CrcUpdatesDifferenceSlice:
		r = TLUpdatesDifferenceSlice{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case CrcUpdatesTooLong:
		r = TLUpdatesTooLong{}

	case CrcUpdateShortMessage:
		r = TLUpdateShortMessage{
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcUpdateShortChatMessage:
		r = TLUpdateShortChatMessage{
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcUpdateShort:
		r = TLUpdateShort{
			m.Object(),
			m.Int(),
		}

	case CrcUpdatesCombined:
		r = TLUpdatesCombined{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcUpdates:
		r = TLUpdates{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case CrcPhotosPhotos:
		r = TLPhotosPhotos{
			m.Vector(),
			m.Vector(),
		}

	case CrcPhotosPhotosSlice:
		r = TLPhotosPhotosSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case CrcPhotosPhoto:
		r = TLPhotosPhoto{
			m.Object(),
			m.Vector(),
		}

	case CrcUploadFile:
		r = TLUploadFile{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case CrcDcOption:
		r = TLDcOption{
			m.Int(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case CrcConfig:
		r = TLConfig{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Vector(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case CrcNearestDc:
		r = TLNearestDc{
			m.String(),
			m.Int(),
			m.Int(),
		}

	case CrcHelpAppUpdate:
		r = TLHelpAppUpdate{
			m.Int(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case CrcHelpNoAppUpdate:
		r = TLHelpNoAppUpdate{}

	case CrcHelpInviteText:
		r = TLHelpInviteText{
			m.String(),
		}

	case CrcMessagesStatedMessagesLinks:
		r = TLMessagesStatedMessagesLinks{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case CrcMessagesStatedMessageLink:
		r = TLMessagesStatedMessageLink{
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case CrcMessagesSentMessageLink:
		r = TLMessagesSentMessageLink{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case CrcInputGeoChat:
		r = TLInputGeoChat{
			m.Int(),
			m.Long(),
		}

	case CrcInputNotifyGeoChatPeer:
		r = TLInputNotifyGeoChatPeer{
			m.Object(),
		}

	case CrcGeoChat:
		r = TLGeoChat{
			m.Int(),
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case CrcGeoChatMessageEmpty:
		r = TLGeoChatMessageEmpty{
			m.Int(),
			m.Int(),
		}

	case CrcGeoChatMessage:
		r = TLGeoChatMessage{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Object(),
		}

	case CrcGeoChatMessageService:
		r = TLGeoChatMessageService{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case CrcGeochatsStatedMessage:
		r = TLGeochatsStatedMessage{
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Int(),
		}

	case CrcGeochatsLocated:
		r = TLGeochatsLocated{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CrcGeochatsMessages:
		r = TLGeochatsMessages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CrcGeochatsMessagesSlice:
		r = TLGeochatsMessagesSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CrcMessageActionGeoChatCreate:
		r = TLMessageActionGeoChatCreate{
			m.String(),
			m.String(),
		}

	case CrcMessageActionGeoChatCheckin:
		r = TLMessageActionGeoChatCheckin{}

	case CrcUpdateNewGeoChatMessage:
		r = TLUpdateNewGeoChatMessage{
			m.Object(),
		}

	case CrcWallPaperSolid:
		r = TLWallPaperSolid{
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case CrcUpdateNewEncryptedMessage:
		r = TLUpdateNewEncryptedMessage{
			m.Object(),
			m.Int(),
		}

	case CrcUpdateEncryptedChatTyping:
		r = TLUpdateEncryptedChatTyping{
			m.Int(),
		}

	case CrcUpdateEncryption:
		r = TLUpdateEncryption{
			m.Object(),
			m.Int(),
		}

	case CrcUpdateEncryptedMessagesRead:
		r = TLUpdateEncryptedMessagesRead{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcEncryptedChatEmpty:
		r = TLEncryptedChatEmpty{
			m.Int(),
		}

	case CrcEncryptedChatWaiting:
		r = TLEncryptedChatWaiting{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcEncryptedChatRequested:
		r = TLEncryptedChatRequested{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case CrcEncryptedChat:
		r = TLEncryptedChat{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
		}

	case CrcEncryptedChatDiscarded:
		r = TLEncryptedChatDiscarded{
			m.Int(),
		}

	case CrcInputEncryptedChat:
		r = TLInputEncryptedChat{
			m.Int(),
			m.Long(),
		}

	case CrcEncryptedFileEmpty:
		r = TLEncryptedFileEmpty{}

	case CrcEncryptedFile:
		r = TLEncryptedFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcInputEncryptedFileEmpty:
		r = TLInputEncryptedFileEmpty{}

	case CrcInputEncryptedFileUploaded:
		r = TLInputEncryptedFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
		}

	case CrcInputEncryptedFile:
		r = TLInputEncryptedFile{
			m.Long(),
			m.Long(),
		}

	case CrcInputEncryptedFileLocation:
		r = TLInputEncryptedFileLocation{
			m.Long(),
			m.Long(),
		}

	case CrcEncryptedMessage:
		r = TLEncryptedMessage{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case CrcEncryptedMessageService:
		r = TLEncryptedMessageService{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case CrcMessagesDhConfigNotModified:
		r = TLMessagesDhConfigNotModified{
			m.StringBytes(),
		}

	case CrcMessagesDhConfig:
		r = TLMessagesDhConfig{
			m.Int(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case CrcMessagesSentEncryptedMessage:
		r = TLMessagesSentEncryptedMessage{
			m.Int(),
		}

	case CrcMessagesSentEncryptedFile:
		r = TLMessagesSentEncryptedFile{
			m.Int(),
			m.Object(),
		}

	case CrcInputFileBig:
		r = TLInputFileBig{
			m.Long(),
			m.Int(),
			m.String(),
		}

	case CrcInputEncryptedFileBigUploaded:
		r = TLInputEncryptedFileBigUploaded{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case CrcUpdateChatParticipantAdd:
		r = TLUpdateChatParticipantAdd{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcUpdateChatParticipantDelete:
		r = TLUpdateChatParticipantDelete{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcUpdateDcOptions:
		r = TLUpdateDcOptions{
			m.Vector(),
		}

	case CrcInputMediaUploadedAudio:
		r = TLInputMediaUploadedAudio{
			m.Object(),
			m.Int(),
			m.String(),
		}

	case CrcInputMediaAudio:
		r = TLInputMediaAudio{
			m.Object(),
		}

	case CrcInputMediaUploadedDocument:
		r = TLInputMediaUploadedDocument{
			m.Object(),
			m.String(),
			m.Vector(),
		}

	case CrcInputMediaUploadedThumbDocument:
		r = TLInputMediaUploadedThumbDocument{
			m.Object(),
			m.Object(),
			m.String(),
			m.Vector(),
		}

	case CrcInputMediaDocument:
		r = TLInputMediaDocument{
			m.Object(),
		}

	case CrcMessageMediaDocument:
		r = TLMessageMediaDocument{
			m.Object(),
		}

	case CrcMessageMediaAudio:
		r = TLMessageMediaAudio{
			m.Object(),
		}

	case CrcInputAudioEmpty:
		r = TLInputAudioEmpty{}

	case CrcInputAudio:
		r = TLInputAudio{
			m.Long(),
			m.Long(),
		}

	case CrcInputDocumentEmpty:
		r = TLInputDocumentEmpty{}

	case CrcInputDocument:
		r = TLInputDocument{
			m.Long(),
			m.Long(),
		}

	case CrcInputAudioFileLocation:
		r = TLInputAudioFileLocation{
			m.Long(),
			m.Long(),
		}

	case CrcInputDocumentFileLocation:
		r = TLInputDocumentFileLocation{
			m.Long(),
			m.Long(),
		}

	case CrcAudioEmpty:
		r = TLAudioEmpty{
			m.Long(),
		}

	case CrcAudio:
		r = TLAudio{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case CrcDocumentEmpty:
		r = TLDocumentEmpty{
			m.Long(),
		}

	case CrcDocument:
		r = TLDocument{
			m.Long(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case CrcHelpSupport:
		r = TLHelpSupport{
			m.String(),
			m.Object(),
		}

	case CrcNotifyPeer:
		r = TLNotifyPeer{
			m.Object(),
		}

	case CrcNotifyUsers:
		r = TLNotifyUsers{}

	case CrcNotifyChats:
		r = TLNotifyChats{}

	case CrcNotifyAll:
		r = TLNotifyAll{}

	case CrcUpdateUserBlocked:
		r = TLUpdateUserBlocked{
			m.Int(),
			m.Object(),
		}

	case CrcUpdateNotifySettings:
		r = TLUpdateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case CrcAuthSentAppCode:
		r = TLAuthSentAppCode{
			m.Object(),
			m.String(),
			m.Int(),
			m.Object(),
		}

	case CrcSendMessageTypingAction:
		r = TLSendMessageTypingAction{}

	case CrcSendMessageCancelAction:
		r = TLSendMessageCancelAction{}

	case CrcSendMessageRecordVideoAction:
		r = TLSendMessageRecordVideoAction{}

	case CrcSendMessageUploadVideoAction:
		r = TLSendMessageUploadVideoAction{}

	case CrcSendMessageRecordAudioAction:
		r = TLSendMessageRecordAudioAction{}

	case CrcSendMessageUploadAudioAction:
		r = TLSendMessageUploadAudioAction{}

	case CrcSendMessageUploadPhotoAction:
		r = TLSendMessageUploadPhotoAction{}

	case CrcSendMessageUploadDocumentAction:
		r = TLSendMessageUploadDocumentAction{}

	case CrcSendMessageGeoLocationAction:
		r = TLSendMessageGeoLocationAction{}

	case CrcSendMessageChooseContactAction:
		r = TLSendMessageChooseContactAction{}

	case CrcContactFound:
		r = TLContactFound{
			m.Int(),
		}

	case CrcContactsFound:
		r = TLContactsFound{
			m.Vector(),
			m.Vector(),
		}

	case CrcUpdateServiceNotification:
		r = TLUpdateServiceNotification{
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case CrcUserStatusRecently:
		r = TLUserStatusRecently{}

	case CrcUserStatusLastWeek:
		r = TLUserStatusLastWeek{}

	case CrcUserStatusLastMonth:
		r = TLUserStatusLastMonth{}

	case CrcUpdatePrivacy:
		r = TLUpdatePrivacy{
			m.Object(),
			m.Vector(),
		}

	case CrcInputPrivacyKeyStatusTimestamp:
		r = TLInputPrivacyKeyStatusTimestamp{}

	case CrcPrivacyKeyStatusTimestamp:
		r = TLPrivacyKeyStatusTimestamp{}

	case CrcInputPrivacyValueAllowContacts:
		r = TLInputPrivacyValueAllowContacts{}

	case CrcInputPrivacyValueAllowAll:
		r = TLInputPrivacyValueAllowAll{}

	case CrcInputPrivacyValueAllowUsers:
		r = TLInputPrivacyValueAllowUsers{
			m.Vector(),
		}

	case CrcInputPrivacyValueDisallowContacts:
		r = TLInputPrivacyValueDisallowContacts{}

	case CrcInputPrivacyValueDisallowAll:
		r = TLInputPrivacyValueDisallowAll{}

	case CrcInputPrivacyValueDisallowUsers:
		r = TLInputPrivacyValueDisallowUsers{
			m.Vector(),
		}

	case CrcPrivacyValueAllowContacts:
		r = TLPrivacyValueAllowContacts{}

	case CrcPrivacyValueAllowAll:
		r = TLPrivacyValueAllowAll{}

	case CrcPrivacyValueAllowUsers:
		r = TLPrivacyValueAllowUsers{
			m.VectorInt(),
		}

	case CrcPrivacyValueDisallowContacts:
		r = TLPrivacyValueDisallowContacts{}

	case CrcPrivacyValueDisallowAll:
		r = TLPrivacyValueDisallowAll{}

	case CrcPrivacyValueDisallowUsers:
		r = TLPrivacyValueDisallowUsers{
			m.VectorInt(),
		}

	case CrcAccountPrivacyRules:
		r = TLAccountPrivacyRules{
			m.Vector(),
			m.Vector(),
		}

	case CrcAccountDaysTTL:
		r = TLAccountDaysTTL{
			m.Int(),
		}

	case CrcAccountSentChangePhoneCode:
		r = TLAccountSentChangePhoneCode{
			m.String(),
			m.Int(),
		}

	case CrcUpdateUserPhone:
		r = TLUpdateUserPhone{
			m.Int(),
			m.String(),
		}

	case CrcDocumentAttributeImageSize:
		r = TLDocumentAttributeImageSize{
			m.Int(),
			m.Int(),
		}

	case CrcDocumentAttributeAnimated:
		r = TLDocumentAttributeAnimated{}

	case CrcDocumentAttributeSticker:
		r = TLDocumentAttributeSticker{}

	case CrcDocumentAttributeVideo:
		r = TLDocumentAttributeVideo{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcDocumentAttributeAudio:
		r = TLDocumentAttributeAudio{
			m.Int(),
		}

	case CrcDocumentAttributeFilename:
		r = TLDocumentAttributeFilename{
			m.String(),
		}

	case CrcMessagesStickersNotModified:
		r = TLMessagesStickersNotModified{}

	case CrcMessagesStickers:
		r = TLMessagesStickers{
			m.String(),
			m.Vector(),
		}

	case CrcStickerPack:
		r = TLStickerPack{
			m.String(),
			m.VectorLong(),
		}

	case CrcMessagesAllStickersNotModified:
		r = TLMessagesAllStickersNotModified{}

	case CrcMessagesAllStickers:
		r = TLMessagesAllStickers{
			m.String(),
			m.Vector(),
			m.Vector(),
		}

	case CrcDisabledFeature:
		r = TLDisabledFeature{
			m.String(),
			m.String(),
		}

	case CrcInvokeAfterMsg:
		r = TLInvokeAfterMsg{
			m.Long(),
			m.Object(),
		}

	case CrcInvokeAfterMsgs:
		r = TLInvokeAfterMsgs{
			m.VectorLong(),
			m.Object(),
		}

	case CrcAuthCheckPhone:
		r = TLAuthCheckPhone{
			m.String(),
		}

	case CrcAuthSendCode:
		r = TLAuthSendCode{
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case CrcAuthSendCall:
		r = TLAuthSendCall{
			m.String(),
			m.String(),
		}

	case CrcAuthSignUp:
		r = TLAuthSignUp{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CrcAuthSignIn:
		r = TLAuthSignIn{
			m.String(),
			m.String(),
			m.String(),
		}

	case CrcAuthLogOut:
		r = TLAuthLogOut{}

	case CrcAuthResetAuthorizations:
		r = TLAuthResetAuthorizations{}

	case CrcAuthSendInvites:
		r = TLAuthSendInvites{
			m.VectorString(),
			m.String(),
		}

	case CrcAuthExportAuthorization:
		r = TLAuthExportAuthorization{
			m.Int(),
		}

	case CrcAuthImportAuthorization:
		r = TLAuthImportAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case CrcAuthBindTempAuthKey:
		r = TLAuthBindTempAuthKey{
			m.Long(),
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case CrcAccountRegisterDevice:
		r = TLAccountRegisterDevice{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.String(),
		}

	case CrcAccountUnregisterDevice:
		r = TLAccountUnregisterDevice{
			m.Int(),
			m.String(),
		}

	case CrcAccountUpdateNotifySettings:
		r = TLAccountUpdateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case CrcAccountGetNotifySettings:
		r = TLAccountGetNotifySettings{
			m.Object(),
		}

	case CrcAccountResetNotifySettings:
		r = TLAccountResetNotifySettings{}

	case CrcAccountUpdateProfile:
		r = TLAccountUpdateProfile{
			m.String(),
			m.String(),
		}

	case CrcAccountUpdateStatus:
		r = TLAccountUpdateStatus{
			m.Object(),
		}

	case CrcAccountGetWallPapers:
		r = TLAccountGetWallPapers{}

	case CrcUsersGetUsers:
		r = TLUsersGetUsers{
			m.Vector(),
		}

	case CrcUsersGetFullUser:
		r = TLUsersGetFullUser{
			m.Object(),
		}

	case CrcContactsGetStatuses:
		r = TLContactsGetStatuses{}

	case CrcContactsGetContacts:
		r = TLContactsGetContacts{
			m.String(),
		}

	case CrcContactsImportContacts:
		r = TLContactsImportContacts{
			m.Vector(),
			m.Object(),
		}

	case CrcContactsGetSuggested:
		r = TLContactsGetSuggested{
			m.Int(),
		}

	case CrcContactsDeleteContact:
		r = TLContactsDeleteContact{
			m.Object(),
		}

	case CrcContactsDeleteContacts:
		r = TLContactsDeleteContacts{
			m.Vector(),
		}

	case CrcContactsBlock:
		r = TLContactsBlock{
			m.Object(),
		}

	case CrcContactsUnblock:
		r = TLContactsUnblock{
			m.Object(),
		}

	case CrcContactsGetBlocked:
		r = TLContactsGetBlocked{
			m.Int(),
			m.Int(),
		}

	case CrcContactsExportCard:
		r = TLContactsExportCard{}

	case CrcContactsImportCard:
		r = TLContactsImportCard{
			m.VectorInt(),
		}

	case CrcMessagesGetMessages:
		r = TLMessagesGetMessages{
			m.VectorInt(),
		}

	case CrcMessagesGetDialogs:
		r = TLMessagesGetDialogs{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcMessagesGetHistory:
		r = TLMessagesGetHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcMessagesSearch:
		r = TLMessagesSearch{
			m.Object(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcMessagesReadHistory:
		r = TLMessagesReadHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case CrcMessagesDeleteHistory:
		r = TLMessagesDeleteHistory{
			m.Object(),
			m.Int(),
		}

	case CrcMessagesDeleteMessages:
		r = TLMessagesDeleteMessages{
			m.VectorInt(),
		}

	case CrcMessagesReceivedMessages:
		r = TLMessagesReceivedMessages{
			m.Int(),
		}

	case CrcMessagesSetTyping:
		r = TLMessagesSetTyping{
			m.Object(),
			m.Object(),
		}

	case CrcMessagesSendMessage:
		r = TLMessagesSendMessage{
			m.Object(),
			m.String(),
			m.Long(),
		}

	case CrcMessagesSendMedia:
		r = TLMessagesSendMedia{
			m.Object(),
			m.Object(),
			m.Long(),
		}

	case CrcMessagesForwardMessages:
		r = TLMessagesForwardMessages{
			m.Object(),
			m.VectorInt(),
		}

	case CrcMessagesGetChats:
		r = TLMessagesGetChats{
			m.VectorInt(),
		}

	case CrcMessagesGetFullChat:
		r = TLMessagesGetFullChat{
			m.Int(),
		}

	case CrcMessagesEditChatTitle:
		r = TLMessagesEditChatTitle{
			m.Int(),
			m.String(),
		}

	case CrcMessagesEditChatPhoto:
		r = TLMessagesEditChatPhoto{
			m.Int(),
			m.Object(),
		}

	case CrcMessagesAddChatUser:
		r = TLMessagesAddChatUser{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case CrcMessagesDeleteChatUser:
		r = TLMessagesDeleteChatUser{
			m.Int(),
			m.Object(),
		}

	case CrcMessagesCreateChat:
		r = TLMessagesCreateChat{
			m.Vector(),
			m.String(),
		}

	case CrcUpdatesGetState:
		r = TLUpdatesGetState{}

	case CrcUpdatesGetDifference:
		r = TLUpdatesGetDifference{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcPhotosUpdateProfilePhoto:
		r = TLPhotosUpdateProfilePhoto{
			m.Object(),
			m.Object(),
		}

	case CrcPhotosUploadProfilePhoto:
		r = TLPhotosUploadProfilePhoto{
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case CrcPhotosDeletePhotos:
		r = TLPhotosDeletePhotos{
			m.Vector(),
		}

	case CrcUploadSaveFilePart:
		r = TLUploadSaveFilePart{
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case CrcUploadGetFile:
		r = TLUploadGetFile{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CrcHelpGetConfig:
		r = TLHelpGetConfig{}

	case CrcHelpGetNearestDc:
		r = TLHelpGetNearestDc{}

	case CrcHelpGetAppUpdate:
		r = TLHelpGetAppUpdate{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CrcHelpSaveAppLog:
		r = TLHelpSaveAppLog{
			m.Vector(),
		}

	case CrcHelpGetInviteText:
		r = TLHelpGetInviteText{
			m.String(),
		}

	case CrcPhotosGetUserPhotos:
		r = TLPhotosGetUserPhotos{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcMessagesForwardMessage:
		r = TLMessagesForwardMessage{
			m.Object(),
			m.Int(),
			m.Long(),
		}

	case CrcMessagesSendBroadcast:
		r = TLMessagesSendBroadcast{
			m.Vector(),
			m.String(),
			m.Object(),
		}

	case CrcGeochatsGetLocated:
		r = TLGeochatsGetLocated{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CrcGeochatsGetRecents:
		r = TLGeochatsGetRecents{
			m.Int(),
			m.Int(),
		}

	case CrcGeochatsCheckin:
		r = TLGeochatsCheckin{
			m.Object(),
		}

	case CrcGeochatsGetFullChat:
		r = TLGeochatsGetFullChat{
			m.Object(),
		}

	case CrcGeochatsEditChatTitle:
		r = TLGeochatsEditChatTitle{
			m.Object(),
			m.String(),
			m.String(),
		}

	case CrcGeochatsEditChatPhoto:
		r = TLGeochatsEditChatPhoto{
			m.Object(),
			m.Object(),
		}

	case CrcGeochatsSearch:
		r = TLGeochatsSearch{
			m.Object(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcGeochatsGetHistory:
		r = TLGeochatsGetHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CrcGeochatsSetTyping:
		r = TLGeochatsSetTyping{
			m.Object(),
			m.Object(),
		}

	case CrcGeochatsSendMessage:
		r = TLGeochatsSendMessage{
			m.Object(),
			m.String(),
			m.Long(),
		}

	case CrcGeochatsSendMedia:
		r = TLGeochatsSendMedia{
			m.Object(),
			m.Object(),
			m.Long(),
		}

	case CrcGeochatsCreateGeoChat:
		r = TLGeochatsCreateGeoChat{
			m.String(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case CrcMessagesGetDhConfig:
		r = TLMessagesGetDhConfig{
			m.Int(),
			m.Int(),
		}

	case CrcMessagesRequestEncryption:
		r = TLMessagesRequestEncryption{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case CrcMessagesAcceptEncryption:
		r = TLMessagesAcceptEncryption{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case CrcMessagesDiscardEncryption:
		r = TLMessagesDiscardEncryption{
			m.Int(),
		}

	case CrcMessagesSetEncryptedTyping:
		r = TLMessagesSetEncryptedTyping{
			m.Object(),
			m.Object(),
		}

	case CrcMessagesReadEncryptedHistory:
		r = TLMessagesReadEncryptedHistory{
			m.Object(),
			m.Int(),
		}

	case CrcMessagesSendEncrypted:
		r = TLMessagesSendEncrypted{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case CrcMessagesSendEncryptedFile:
		r = TLMessagesSendEncryptedFile{
			m.Object(),
			m.Long(),
			m.StringBytes(),
			m.Object(),
		}

	case CrcMessagesSendEncryptedService:
		r = TLMessagesSendEncryptedService{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case CrcMessagesReceivedQueue:
		r = TLMessagesReceivedQueue{
			m.Int(),
		}

	case CrcUploadSaveBigFilePart:
		r = TLUploadSaveBigFilePart{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case CrcInitConnection:
		r = TLInitConnection{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
		}

	case CrcHelpGetSupport:
		r = TLHelpGetSupport{}

	case CrcAuthSendSms:
		r = TLAuthSendSms{
			m.String(),
			m.String(),
		}

	case CrcMessagesReadMessageContents:
		r = TLMessagesReadMessageContents{
			m.VectorInt(),
		}

	case CrcAccountCheckUsername:
		r = TLAccountCheckUsername{
			m.String(),
		}

	case CrcAccountUpdateUsername:
		r = TLAccountUpdateUsername{
			m.String(),
		}

	case CrcContactsSearch:
		r = TLContactsSearch{
			m.String(),
			m.Int(),
		}

	case CrcAccountGetPrivacy:
		r = TLAccountGetPrivacy{
			m.Object(),
		}

	case CrcAccountSetPrivacy:
		r = TLAccountSetPrivacy{
			m.Object(),
			m.Vector(),
		}

	case CrcAccountDeleteAccount:
		r = TLAccountDeleteAccount{
			m.String(),
		}

	case CrcAccountGetAccountTTL:
		r = TLAccountGetAccountTTL{}

	case CrcAccountSetAccountTTL:
		r = TLAccountSetAccountTTL{
			m.Object(),
		}

	case CrcInvokeWithLayer:
		r = TLInvokeWithLayer{
			m.Int(),
			m.Object(),
		}

	case CrcContactsResolveUsername:
		r = TLContactsResolveUsername{
			m.String(),
		}

	case CrcAccountSendChangePhoneCode:
		r = TLAccountSendChangePhoneCode{
			m.String(),
		}

	case CrcAccountChangePhone:
		r = TLAccountChangePhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case CrcMessagesGetStickers:
		r = TLMessagesGetStickers{
			m.String(),
			m.String(),
		}

	case CrcMessagesGetAllStickers:
		r = TLMessagesGetAllStickers{
			m.String(),
		}

	case CrcAccountUpdateDeviceLocked:
		r = TLAccountUpdateDeviceLocked{
			m.Int(),
		}

	default:
		m.err = fmt.Errorf("unknown constructor: \u002508x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}
