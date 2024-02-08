package Estructuras

type MBR struct {
	MBR_SIZE int64
	MBR_DATE [16]byte
	MBR_ID   int64
}

func NewMBR() MBR {

	return MBR{
		MBR_SIZE: 0,
		MBR_DATE: [16]byte{},
		MBR_ID:   0,
	}

}
