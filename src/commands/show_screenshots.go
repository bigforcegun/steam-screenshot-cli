package commands

import "fmt"

func ShowScreenshots(userID string) {
	/*
		ну карочи блеть

		зашли сюда - https://steamcommunity.com/id/bigforcegun/screenshots/#scrollTop=0
		получили скринты
		https://steamcommunity.com/sharedfiles/filedetails/?id=3095228393 - зашли по ID файла
		получили файл и мета инфу
		сохранили
		...
		профит

		вариант 2
		достать по API список игор
		https://steamcommunity.com/id/bigforcegun/screenshots/?appid=620#scrollTop=25
		по appid ходить по игорам и выкачивать скрины

		API не дает инфу по скринам, только по установленым игорам
	*/
	fmt.Println("go go to screenshots", userID)
}
