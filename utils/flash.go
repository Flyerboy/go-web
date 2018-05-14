package utils

const (
	FLASH_UNKNOW = 0
	FLASH_SUCCESS = 1
	FLASH_ERROR = 2
	FLASH_WARNING = 3
)

type FlashInterface interface {
	Get() (level int, message string)
	Destroy()
	Info(level int, message string)
	Success(message string)
	Error(message string)
	Warning(message string)
	Write()
}


type Flash struct {
	Level int
	Message string
}

func (this *Flash) Get() (level int, message string) {
	level = this.Level
	message = this.Message
	this.Destroy()
	return level, message
}

func (this *Flash) Destroy() {
	this.Level = FLASH_UNKNOW
	this.Message = ""
}

func (this *Flash) Info(level int, message string)  {
	this.Message = message
	this.Level = level
	this.Write()
}

func (this *Flash) Success(message string)  {
	this.Message = message
	this.Level = FLASH_SUCCESS
	this.Write()
}

func (this *Flash) Error(message string)  {
	this.Message = message
	this.Level = FLASH_ERROR
	this.Write()
}

func (this *Flash) Warning(message string)  {
	this.Message = message
	this.Level = FLASH_WARNING
	this.Write()
}

func (this *Flash) Write() {

}

