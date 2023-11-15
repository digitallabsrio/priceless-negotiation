package min

type Miner interface {
	Len() int
	ElemIx(ix int) interface{}
	Less(i, j int) bool
}

func Min(data Miner) interface{}  {
	
}

type IntArray []int
func (p IntArray) Len() int           		  {}
func (p IntArray) ElemIx(ix int) interface{}  {}
func (p IntArray) Less(i, j int) bool 		  {}

type StringArray []string
func (p StringArray) Len() int              	 {}
func (p StringArray) ElemIx(ix int) interface{}  {}
func (p StringArray) Less(i, j int) bool    	 {}