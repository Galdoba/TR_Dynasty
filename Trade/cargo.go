package Trade

type Cargo struct {
	code            string
	sourceWorldUPP  string
	aquisitionPrice int
	volume          int
}

/*
код груза:
A B C D E F G container type container size container mass atmospheric range temperature range humidity range gravity range
H I J K L M N O cargo type number of items mass of each item atmospheric range temperature range humidity range gravity range EM spectrum range
CONTAINER:
container type				0-4
container size				0-9
container mass 				0-9
atmospheric range 			0-9
temperature range 			X 0-F Y
humidity range 				0-A
gravity range 				0-X
CARGO:
cargo type 					0-E
number of items 			0-B
mass of each item 			0-9
atmospheric range 			0-9
temperature range			X 0-F Y
humidity range				0-A
gravity range 				0-X
EM spectrum range			0-8
*/
