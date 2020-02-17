// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package gui

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"ar": &dictionary{index: arIndex, data: arData},
		"en": &dictionary{index: enIndex, data: enData},
		"es": &dictionary{index: esIndex, data: esData},
		"sv": &dictionary{index: svIndex, data: svData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%sMeeting ID: %s":                    116,
	"A valid port is between 1 and 65535": 98,
	"Accept":                              28,
	"Allow the host to automatically join a newly created meeting": 29,
	"An error occurred\n\n%s":                                      14,
	"Are you sure you want to do this action?":                     30,
	"Are you sure you want to end this meeting?":                   31,
	"Are you sure you want to leave this meeting?":                 32,
	"Automatically join a meeting":                                 34,
	"Automatically join this meeting":                              33,
	"Automatically join this meeting when starting it":             35,
	"Be very careful. This information is sensitive and could potentially contain very private information. Only turn on these settings if you absolutely need it for debugging.": 36,
	"Browse": 37,
	"By clicking Yes, this meeting will end.":       38,
	"By clicking Yes, you will leave this meeting.": 39,
	"Cancel": 24,
	"Check this option to automatically join every meeting you host": 49,
	"Choose your email service to send invitation":                   50,
	"Client binary location":                                         40,
	"Configuration settings will be lost in the next session":        41,
	"Configure master password":                                      42,
	"Configured Mumble port is not valid: %s":                        3,
	"Confirmation":                   43,
	"Connecting, please wait...":     44,
	"Continue":                       45,
	"Copy Invitation":                46,
	"Copy Meeting ID":                47,
	"Copy URL":                       48,
	"Debugging":                      51,
	"Default Email":                  52,
	"Encrypt the configuration file": 53,
	"End this meeting":               55,
	"End this meeting for all":       56,
	"Error":                          0,
	"Finish":                         54,
	"General":                        57,
	"Gmail":                          58,
	"Host a new meeting":             59,
	"Host meeting":                   61,
	"Hosting":                        60,
	"If you backup the configuration file, we will reset the settings and continue normally. If the configuration file is encrypted, then we will ask you for a password to encrypt the new settings file.": 62,
	"If you disable this option, anyone could read your configuration settings":          22,
	"If you set this option to a file name, low level information will be logged there.": 63,
	"Invalid configuration file":                64,
	"Invalid password. Please, try again.":      65,
	"Invite others":                             66,
	"Join":                                      67,
	"Join Wahay Meeting":                        9,
	"Join a meeting":                            68,
	"Join meeting":                              69,
	"Join the meeting":                          70,
	"Join this meeting":                         71,
	"Keep configuration file when Wahay closes": 72,
	"Leave":              73,
	"Leave this meeting": 74,
	"Log debug info":     75,
	"Log debug output to the selected log file. If no file is selected then the log output will be written to the default log file.": 76,
	"Master password":                77,
	"Meeting ID":                     78,
	"Meeting ID:":                    79,
	"Meeting password":               80,
	"Mumble":                         81,
	"No, cancel":                     82,
	"Now you are hosting a meeting.": 83,
	"Open":                           25,
	"Open file":                      23,
	"Outlook":                        84,
	"Password":                       85,
	"Please enter the master password for the configuration file.": 86,
	"Please join the Wahay meeting with the following details:":    115,
	"Port":                               87,
	"Port out of range":                  88,
	"Raw log file":                       89,
	"Repeat the password":                90,
	"Save changes":                       91,
	"Security":                           92,
	"Settings":                           93,
	"Show":                               94,
	"Something went wrong: %s":           1,
	"Specify a password for the meeting": 95,
	"Start Meeting":                      11,
	"Start Meeting & Join":               10,
	"Start meeting":                      96,
	"The Meeting ID cannot be blank":     13,
	"The Mumble process is down":         12,
	"The error message":                  97,
	"The invitation email has been copied to the clipboard": 8,
	"The meeting ID has been copied to the clipboard":       7,
	"The meeting ID is invalid":                             16,
	"The meeting can't be closed: %s":                       4,
	"The onion service can't be deleted: %s":                6,
	"This action cannot be undone":                          99,
	"Toggle password visibility":                            100,
	"Type the Meeting ID (normally a .onion address)":       101,
	"Type the password":                                     102,
	"Type the password to join the meeting":                 103,
	"Type your preferred screen name":                       104,
	"Type your screen name":                                 105,
	"Username":                                              106,
	"Wahay is ready to use":                                 107,
	"We have detected that the configuration file is invalid or corrupted. Do you want to make a copy (backup) of it and continue?": 108,
	"We've found errors": 27,
	"Welcome":            109,
	"When this option is checked, the configuration settings will be stored in the device.": 110,
	"Yahoo Mail":                     111,
	"Yes, back it up &amp; continue": 112,
	"Yes, confirm":                   113,
	"You will not be asked for this password again until you restart Wahay.": 114,
	"enter a password at least 6 characters long":                            21,
	"enter the password confirmation":                                        19,
	"internal Tor instance has already been closed":                          5,
	"passwords do not match":                                                 20,
	"please enter a valid password":                                          18,
	"the Mumble client can not be used because: %s":                          17,
	"the provided meeting ID is invalid: \n\n%s":                             15,
	"tor can't be used":                                                      26,
	"we couldn't start the meeting":                                          2,
}

var arIndex = []uint32{ // 118 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000d, 0x0000001a, 0x00000027,
	0x00000034, 0x00000041, 0x0000004e, 0x0000005b,
	0x00000068, 0x00000075, 0x00000082, 0x0000008f,
	0x0000009c, 0x000000a9, 0x000000b6, 0x000000c3,
	0x000000d0, 0x000000dd, 0x000000ea, 0x000000f7,
	0x00000104, 0x00000111, 0x0000011e, 0x0000012b,
	0x00000138, 0x00000145, 0x00000152, 0x0000015f,
	0x0000016c, 0x00000179, 0x00000186, 0x00000193,
	// Entry 20 - 3F
	0x000001a0, 0x000001ad, 0x000001ba, 0x000001c7,
	0x000001d4, 0x000001e1, 0x000001ee, 0x000001fb,
	0x00000208, 0x00000215, 0x00000222, 0x0000022f,
	0x0000023c, 0x0000026c, 0x00000279, 0x00000286,
	0x00000293, 0x000002a0, 0x000002ad, 0x000002ba,
	0x000002c7, 0x000002d4, 0x000002e1, 0x000002ee,
	0x000002fb, 0x00000308, 0x00000315, 0x00000322,
	0x0000032f, 0x0000033c, 0x00000349, 0x00000356,
	// Entry 40 - 5F
	0x00000363, 0x00000370, 0x0000037d, 0x0000038a,
	0x00000397, 0x000003a4, 0x000003b1, 0x000003be,
	0x000003cb, 0x000003d8, 0x000003e5, 0x000003f2,
	0x000003ff, 0x0000040c, 0x00000419, 0x00000426,
	0x00000433, 0x00000440, 0x0000044d, 0x0000045a,
	0x00000467, 0x00000474, 0x00000481, 0x0000048e,
	0x0000049b, 0x000004a8, 0x000004b5, 0x000004c2,
	0x000004cf, 0x000004dc, 0x000004e9, 0x000004f6,
	// Entry 60 - 7F
	0x00000503, 0x00000510, 0x0000051d, 0x0000052a,
	0x00000537, 0x00000544, 0x00000551, 0x0000055e,
	0x0000056b, 0x00000578, 0x00000585, 0x00000592,
	0x0000059f, 0x000005ac, 0x000005b9, 0x000005c6,
	0x000005d3, 0x000005e0, 0x000005ed, 0x000005fa,
	0x000005fa, 0x000005fa,
} // Size: 496 bytes

const arData string = "" + // Size: 1530 bytes
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02...الاتصال ا" +
	"لرجاء الانتظار\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANS" +
	"LATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME"

var enIndex = []uint32{ // 118 elements
	// Entry 0 - 1F
	0x00000000, 0x00000006, 0x00000022, 0x00000040,
	0x0000006b, 0x0000008e, 0x000000bc, 0x000000e6,
	0x00000116, 0x0000014c, 0x0000015f, 0x00000174,
	0x00000182, 0x0000019d, 0x000001bc, 0x000001d5,
	0x00000201, 0x0000021b, 0x0000024c, 0x0000026a,
	0x0000028a, 0x000002a1, 0x000002cd, 0x00000317,
	0x00000321, 0x00000328, 0x0000032d, 0x0000033f,
	0x00000352, 0x00000359, 0x00000396, 0x000003bf,
	// Entry 20 - 3F
	0x000003ea, 0x00000417, 0x00000437, 0x00000454,
	0x00000485, 0x00000531, 0x00000538, 0x00000560,
	0x0000058e, 0x000005a5, 0x000005dd, 0x000005f7,
	0x00000604, 0x0000061f, 0x00000628, 0x00000638,
	0x00000648, 0x00000651, 0x00000690, 0x000006bd,
	0x000006c7, 0x000006d5, 0x000006f4, 0x000006fb,
	0x0000070c, 0x00000725, 0x0000072d, 0x00000733,
	0x00000746, 0x0000074e, 0x0000075b, 0x00000821,
	// Entry 40 - 5F
	0x00000874, 0x0000088f, 0x000008b4, 0x000008c2,
	0x000008c7, 0x000008d6, 0x000008e3, 0x000008f4,
	0x00000906, 0x00000930, 0x00000936, 0x00000949,
	0x00000958, 0x000009d7, 0x000009e7, 0x000009f2,
	0x000009fe, 0x00000a0f, 0x00000a16, 0x00000a21,
	0x00000a40, 0x00000a48, 0x00000a51, 0x00000a8e,
	0x00000a93, 0x00000aa5, 0x00000ab2, 0x00000ac6,
	0x00000ad3, 0x00000adc, 0x00000ae5, 0x00000aea,
	// Entry 60 - 7F
	0x00000b0d, 0x00000b1b, 0x00000b2d, 0x00000b51,
	0x00000b6e, 0x00000b89, 0x00000bb9, 0x00000bcb,
	0x00000bf1, 0x00000c11, 0x00000c27, 0x00000c30,
	0x00000c46, 0x00000cc4, 0x00000ccc, 0x00000d22,
	0x00000d2d, 0x00000d4c, 0x00000d59, 0x00000da0,
	0x00000dda, 0x00000df1,
} // Size: 496 bytes

const enData string = "" + // Size: 3569 bytes
	"\x02Error\x02Something went wrong: %[1]s\x02we couldn't start the meetin" +
	"g\x02Configured Mumble port is not valid: %[1]s\x02The meeting can't be " +
	"closed: %[1]s\x02internal Tor instance has already been closed\x02The on" +
	"ion service can't be deleted: %[1]s\x02The meeting ID has been copied to" +
	" the clipboard\x02The invitation email has been copied to the clipboard" +
	"\x02Join Wahay Meeting\x02Start Meeting & Join\x02Start Meeting\x02The M" +
	"umble process is down\x02The Meeting ID cannot be blank\x02An error occu" +
	"rred\x0a\x0a%[1]s\x02the provided meeting ID is invalid: \x0a\x0a%[1]s" +
	"\x02The meeting ID is invalid\x02the Mumble client can not be used becau" +
	"se: %[1]s\x02please enter a valid password\x02enter the password confirm" +
	"ation\x02passwords do not match\x02enter a password at least 6 character" +
	"s long\x02If you disable this option, anyone could read your configurati" +
	"on settings\x02Open file\x02Cancel\x02Open\x02tor can't be used\x02We've" +
	" found errors\x02Accept\x02Allow the host to automatically join a newly " +
	"created meeting\x02Are you sure you want to do this action?\x02Are you s" +
	"ure you want to end this meeting?\x02Are you sure you want to leave this" +
	" meeting?\x02Automatically join this meeting\x02Automatically join a mee" +
	"ting\x02Automatically join this meeting when starting it\x02Be very care" +
	"ful. This information is sensitive and could potentially contain very pr" +
	"ivate information. Only turn on these settings if you absolutely need it" +
	" for debugging.\x02Browse\x02By clicking Yes, this meeting will end.\x02" +
	"By clicking Yes, you will leave this meeting.\x02Client binary location" +
	"\x02Configuration settings will be lost in the next session\x02Configure" +
	" master password\x02Confirmation\x02Connecting, please wait...\x02Contin" +
	"ue\x02Copy Invitation\x02Copy Meeting ID\x02Copy URL\x02Check this optio" +
	"n to automatically join every meeting you host\x02Choose your email serv" +
	"ice to send invitation\x02Debugging\x02Default Email\x02Encrypt the conf" +
	"iguration file\x02Finish\x02End this meeting\x02End this meeting for all" +
	"\x02General\x02Gmail\x02Host a new meeting\x02Hosting\x02Host meeting" +
	"\x02If you backup the configuration file, we will reset the settings and" +
	" continue normally. If the configuration file is encrypted, then we will" +
	" ask you for a password to encrypt the new settings file.\x02If you set " +
	"this option to a file name, low level information will be logged there." +
	"\x02Invalid configuration file\x02Invalid password. Please, try again." +
	"\x02Invite others\x02Join\x02Join a meeting\x02Join meeting\x02Join the " +
	"meeting\x02Join this meeting\x02Keep configuration file when Wahay close" +
	"s\x02Leave\x02Leave this meeting\x02Log debug info\x02Log debug output t" +
	"o the selected log file. If no file is selected then the log output will" +
	" be written to the default log file.\x02Master password\x02Meeting ID" +
	"\x02Meeting ID:\x02Meeting password\x02Mumble\x02No, cancel\x02Now you a" +
	"re hosting a meeting.\x02Outlook\x02Password\x02Please enter the master " +
	"password for the configuration file.\x02Port\x02Port out of range\x02Raw" +
	" log file\x02Repeat the password\x02Save changes\x02Security\x02Settings" +
	"\x02Show\x02Specify a password for the meeting\x02Start meeting\x02The e" +
	"rror message\x02A valid port is between 1 and 65535\x02This action canno" +
	"t be undone\x02Toggle password visibility\x02Type the Meeting ID (normal" +
	"ly a .onion address)\x02Type the password\x02Type the password to join t" +
	"he meeting\x02Type your preferred screen name\x02Type your screen name" +
	"\x02Username\x02Wahay is ready to use\x02We have detected that the confi" +
	"guration file is invalid or corrupted. Do you want to make a copy (backu" +
	"p) of it and continue?\x02Welcome\x02When this option is checked, the co" +
	"nfiguration settings will be stored in the device.\x02Yahoo Mail\x02Yes," +
	" back it up &amp; continue\x02Yes, confirm\x02You will not be asked for " +
	"this password again until you restart Wahay.\x02Please join the Wahay me" +
	"eting with the following details:\x02%[1]sMeeting ID: %[2]s"

var esIndex = []uint32{ // 118 elements
	// Entry 0 - 1F
	0x00000000, 0x00000006, 0x0000001d, 0x0000003d,
	0x00000073, 0x00000098, 0x000000ae, 0x000000db,
	0x0000010f, 0x00000148, 0x00000171, 0x0000018c,
	0x0000019e, 0x000001bd, 0x000001e6, 0x000001ff,
	0x00000233, 0x00000255, 0x00000286, 0x000002b4,
	0x000002e3, 0x00000301, 0x00000335, 0x00000393,
	0x000003a1, 0x000003aa, 0x000003b0, 0x000003c5,
	0x000003e1, 0x000003e9, 0x00000441, 0x00000473,
	// Entry 20 - 3F
	0x000004a9, 0x000004db, 0x00000503, 0x00000538,
	0x0000056d, 0x00000621, 0x0000062a, 0x0000065b,
	0x0000068b, 0x000006a9, 0x000006eb, 0x0000070a,
	0x00000718, 0x00000738, 0x00000742, 0x00000755,
	0x0000076b, 0x00000776, 0x000007dc, 0x00000821,
	0x0000082d, 0x00000842, 0x00000866, 0x0000086f,
	0x00000885, 0x000008a6, 0x000008ae, 0x000008b4,
	0x000008d1, 0x000008d9, 0x000008ed, 0x000009e9,
	// Entry 40 - 5F
	0x00000a50, 0x00000a74, 0x00000a9f, 0x00000aaf,
	0x00000ab6, 0x00000acc, 0x00000ae1, 0x00000af6,
	0x00000b0d, 0x00000b43, 0x00000b49, 0x00000b60,
	0x00000b86, 0x00000c44, 0x00000c58, 0x00000c6a,
	0x00000c7d, 0x00000c98, 0x00000c9f, 0x00000cac,
	0x00000cd3, 0x00000cdb, 0x00000ce7, 0x00000d29,
	0x00000d30, 0x00000d46, 0x00000d5b, 0x00000d71,
	0x00000d81, 0x00000d8b, 0x00000d9b, 0x00000da3,
	// Entry 60 - 7F
	0x00000dd0, 0x00000de4, 0x00000df8, 0x00000e2b,
	0x00000e4d, 0x00000e71, 0x00000eb2, 0x00000ec9,
	0x00000efa, 0x00000f21, 0x00000f3e, 0x00000f50,
	0x00000f6e, 0x00000ff0, 0x00000ffb, 0x00001050,
	0x00001060, 0x0000107d, 0x0000108b, 0x000010d4,
	0x00001123, 0x00001123,
} // Size: 496 bytes

const esData string = "" + // Size: 4387 bytes
	"\x02Error\x02Algo salió mal: %[1]s\x02no se pudo comenzar la reunión\x02" +
	"El puerto configurado para Mumble es inválido: %[1]s\x02La reunión no se" +
	" pudo cerrar: %[1]s\x02tor no está definido\x02El servicio Onion no se p" +
	"udo eliminar: %[1]s\x02El ID de la reunion ha sido copiado al portapapel" +
	"es\x02El correo de invitación ha sido copiado al portapapeles\x02Unirse " +
	"a una reunión a travéz de Wahay\x02Comenzar reunión y Unirse\x02Comenzar" +
	" reunión\x02El proceso Mumble está caído\x02El ID de la reunión no puede" +
	" ser vacío\x02Ocurrió un error\x0a\x0a%[1]s\x02el ID de la reunión provi" +
	"sto es inválido: \x0a\x0a%[1]s\x02El ID de la reunión es inválido\x02El " +
	"cliente Mumble no se puede usar porque: %[1]s\x02por favor especifique u" +
	"na contraseña válida\x02especifique la confirmación de la contraseña\x02" +
	"las contraseñas no coinciden\x02especifique una contraseña de mínimo 6 c" +
	"aracteres\x02Si deshabilita esta opción, cualquier persona podría leer l" +
	"os parámetros de configuración\x02Abrir archivo\x02Cancelar\x02Abrir\x02" +
	"tor no se puede usar\x02Encontramos algunos errores\x02Aceptar\x02Permit" +
	"a que el organizador se una automáticamente a una reunión cuando cree un" +
	"a nueva\x02¿Está seguro de que quieres hacer esta acción?\x02¿Está segur" +
	"o de que quieres terminar esta reunión?\x02¿Está seguro de que quiere de" +
	"jar esta reunión?\x02Unirse automáticamente a esta reunión\x02Únase auto" +
	"máticamente a esta reunión al iniciarla\x02Únase automáticamente a esta " +
	"reunión al iniciarla\x02Ten mucho cuidado. Esta información es confidenc" +
	"ial y podría contener información muy privada. Solo cambia esta configur" +
	"ación si la necesita absolutamente para la depuración.\x02Examinar\x02Al" +
	" hacer clic en Sí, esta reunión finalizará.\x02Al hacer clic en Sí, sald" +
	"rá de esta reunión.\x02Ubicación del binario Mumble\x02Los ajustes de co" +
	"nfiguración se perderán en la próxima sesión\x02Configurar contraseña ma" +
	"estra\x02Confirmación\x02Conectando, espere por favor...\x02Continuar" +
	"\x02Copiar invitación\x02Copiar ID de reunión\x02Copiar URL\x02Marque es" +
	"ta opción para unirse automáticamente a cada reunión creada en la secció" +
	"n del anfitrión\x02Elija su servicio de correo electrónico para enviar l" +
	"a invitación.\x02Depuración\x02Email predeterminado\x02Cifrar el archivo" +
	" de configuración\x02Terminar\x02Termina esta reunión\x02Termina esta re" +
	"unión para todos\x02General\x02Gmail\x02Organizar una nueva reunión\x02H" +
	"osting\x02Reunión de acogida\x02Si realiza una copia de seguridad del ar" +
	"chivo de configuración, restableceremos la configuración y continuaremos" +
	" normalmente. Si el archivo de configuración está cifrado, le pediremos " +
	"una contraseña para cifrar el nuevo archivo de configuración.\x02Si esta" +
	"blece esta opción en un nombre de archivo, la información de bajo nivel " +
	"se registrará allí.\x02Archivo de configuración inválido\x02Contraseña i" +
	"nvalida. Inténtalo de nuevo.\x02Invitar a otros\x02Unirse\x02Unirse a un" +
	"a reunión\x02Unirse a la reunión\x02Únete a la reunión\x02Únete a esta r" +
	"eunión\x02Mantener el archivo de configuración al cerrar Wahay\x02Salir" +
	"\x02Salir de esta reunión\x02Registrar información de depuración\x02Regi" +
	"stre la salida de depuración en el archivo de registro seleccionado. Si " +
	"no se selecciona ningún archivo, la salida del registro se escribirá en " +
	"el archivo de registro predeterminado.\x02Contraseña maestra\x02ID de la" +
	" reunión\x02ID de la reunión:\x02Contraseña de la reunión\x02Mumble\x02N" +
	"o, cancelar\x02Ahora estás organizando una reunión.\x02Outlook\x02Contra" +
	"seña\x02Ingrese la contraseña maestra para el archivo de configuración." +
	"\x02Puerto\x02Puerto fuera de rango\x02Archivo de registros\x02Repita la" +
	" contraseña\x02Guardar cambios\x02Seguridad\x02Configuraciones\x02Mostra" +
	"r\x02Especifique una contraseña para la reunión\x02Comience a reunirse" +
	"\x02El mensaje de error\x02El rango de puertos válidos está entre 1 y 65" +
	"535\x02Esta acción no se puede deshacer\x02Alternar visibilidad de contr" +
	"aseña\x02Escriba la ID de la reunión (normalmente una dirección .onion)" +
	"\x02Escribe la contraseña\x02Escriba la contraseña para unirse a la reun" +
	"ión\x02Escriba su nombre de usuario preferido\x02Escriba su nombre de us" +
	"uario\x02Nombre de usuario\x02Wahay está listo para usarse\x02Hemos dete" +
	"ctado que el archivo de configuración no es válido o está dañado. ¿Desea" +
	" hacer una copia de seguridad y continuar?\x02Bienvenido\x02Cuando esta " +
	"opción está marcada, la configuración se guardará en el dispositivo.\x02" +
	"Correo de Yahoo\x02Sí, respaldarlo y continuar\x02Si, confirmar\x02No se" +
	" le volverá a solicitar esta contraseña hasta que reinicie Wahay.\x02Por" +
	" favor únete a la reunión a travéz de Wahay con los siguientes detalles:"

var svIndex = []uint32{ // 118 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000d, 0x0000001a, 0x00000027,
	0x00000034, 0x00000041, 0x0000004e, 0x0000005b,
	0x00000068, 0x00000075, 0x00000082, 0x0000008f,
	0x0000009c, 0x000000a9, 0x000000b6, 0x000000c3,
	0x000000d0, 0x000000dd, 0x000000ea, 0x000000f7,
	0x00000104, 0x00000111, 0x0000011e, 0x0000012b,
	0x00000138, 0x00000145, 0x00000152, 0x0000015f,
	0x0000016c, 0x00000179, 0x00000186, 0x00000193,
	// Entry 20 - 3F
	0x000001a0, 0x000001ad, 0x000001ba, 0x000001c7,
	0x000001d4, 0x000001e1, 0x000001ee, 0x000001fb,
	0x00000208, 0x00000215, 0x00000222, 0x0000022f,
	0x0000023c, 0x00000258, 0x00000265, 0x00000272,
	0x0000027f, 0x0000028c, 0x00000299, 0x000002a6,
	0x000002b3, 0x000002c0, 0x000002cd, 0x000002da,
	0x000002e7, 0x000002f4, 0x00000301, 0x0000030e,
	0x0000031b, 0x00000328, 0x00000335, 0x00000342,
	// Entry 40 - 5F
	0x0000034f, 0x0000035c, 0x00000369, 0x00000376,
	0x00000383, 0x00000390, 0x0000039d, 0x000003aa,
	0x000003b7, 0x000003c4, 0x000003d1, 0x000003de,
	0x000003eb, 0x000003f8, 0x00000405, 0x00000412,
	0x0000041f, 0x0000042c, 0x00000439, 0x00000446,
	0x00000453, 0x00000460, 0x0000046d, 0x0000047a,
	0x00000487, 0x00000494, 0x000004a1, 0x000004ae,
	0x000004bb, 0x000004c8, 0x000004d5, 0x000004e2,
	// Entry 60 - 7F
	0x000004ef, 0x000004fc, 0x00000509, 0x00000516,
	0x00000523, 0x00000530, 0x0000053d, 0x0000054a,
	0x00000557, 0x00000564, 0x00000571, 0x0000057e,
	0x0000058b, 0x00000598, 0x000005a5, 0x000005b2,
	0x000005bf, 0x000005cc, 0x000005d9, 0x000005e6,
	0x000005f3, 0x000005f3,
} // Size: 496 bytes

const svData string = "" + // Size: 1523 bytes
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02Ansluter, va" +
	"r god vänta...\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSL" +
	"ATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME" +
	"\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRAN" +
	"SLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME\x02TRANSLATE ME"

	// Total table size 12993 bytes (12KiB); checksum: D54FD010
