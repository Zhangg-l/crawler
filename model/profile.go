package model

// ss := `
// <div class="m-btn purple" data-v-8b1eac0c>未婚</div>
// <div class="m-btn pink" data-v-8b1eac0c>籍贯:辽宁鞍山</div>
// `

type Profile struct {
	// 基本信息
	Age           string
	Name          string
	Height        string
	Weight        string
	Eduction      string
	Marriage      string
	Income        string
	Cnostellation string //星座
	WorkPlace     string

	//其他信息
	Car         string
	House       string
	Figure      string
	NativePlace string
}
