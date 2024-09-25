package telegram

const msgHelp = `
Просто пришли мне ссылку вида http:\\... и я сохраню её в твой список. 
Потом напиши мне /rnd, и я отправлю тебе случайню ссылку из списка
ВНИМАНИЕ. После того, как вы получите ссылку - ссылка удалится из вашего списка

Чтобы посмотреть количество сохраненных страниц - /count`

const msgHello = `Пр-пр-пр-пр привет! (типа звук, как заводиться старая карга)
Это телеграм-бот, который поможет сохранять и выдавать случайную ссылку для чтения
Напиши /help для доп. информации.`

const (
	msgUnknownCommand = "Неизвестная команда. Прочитай блять инстуркцию!☠️"
	msgNoSavedPage    = "Ты нихуя не сохранил 👀"
	msgSave           = "Готово ✅"
	msgAlreadyExists  = "Ты это уже сохранял. Найди новое 🙆"
)
