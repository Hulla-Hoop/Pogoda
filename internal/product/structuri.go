package product

type Jandex struct {
	Now      int64         `json:"now"`    //	Время сервера в формате Unixtime.
	Nowutc   string        `json:"now_dt"` // Время сервера в UTC.
	Info     `json:"info"` //Объект информации о населенном пункте.
	Fact     `json:"fact"`
	Forecast `json:"forecast"`
}

type Info struct {
	Lat               float64         `json:"lat"` //Широта (в градусах).
	Lon               float64         `json:"lon"` //Долгота (в градусах).
	TzInfo            `json:"tzinfo"` //Информация о часовом поясе
	NormadavleniaMM   int64           `json:"def_pressure_mm"` //Норма давления для данной координаты (в мм рт. ст.).
	NormadavleniaGECT int64           `json:"def_pressure_pa"` //Норма давления для данной координаты (в гектопаскалях).
	Url               string          `json:"url"`             //Страница населенного пункта на сайте Яндекс.Погода
}

type TzInfo struct {
	Offset int64  `json:"offset"` //Часовой пояс в секундах от UTC.
	Name   string `json:"name"`   //Название часового пояса.
	Abbr   string `json:"abbr"`   //	Сокращенное название часового пояса.
	Dst    bool   `json:"dst"`    //	Признак летнего времени.
}

type Fact struct {
	Temp      int    `json:"temp"`       // Температура (°C)
	FeelsLike int    `json:"feels_like"` // Ощущаемая температура (°C)
	TempWater int    `json:"temp_water"` // Температура воды (°C). Параметр возвращается для населенных пунктов, где данная информация актуальна
	Icon      string `json:"icon"`       // Код иконки погоды, иконка доступна по адресу https://yastatic.net/weather/i/icons/funky/dark/<значение из поля icon>.svg
	Condition string `json:"condition"`  // Код расшифровки погодного описания. Возможные значения:
	//clear — ясно
	//partly-cloudy — малооблачно
	//cloudy — облачно с прояснениями
	//overcast — пасмурно
	//drizzle — морось
	//light-rain — небольшой дождь
	//rain — дождь
	//moderate-rain — умеренно сильный дождь
	//heavy-rain — сильный дождь
	//continuous-heavy-rain — длительный сильный дождь
	//showers — ливень
	//wet-snow — дождь со снегом
	//light-snow — небольшой снег
	//snow — снег
	//snow-showers — снегопад
	//hail — град
	//thunderstorm — гроза
	//thunderstorm-with-rain — дождь с грозой
	//thunderstorm-with-hail — гроза с градом
	WindSpeed float64 `json:"wind_speed"` // Скорость ветра (в м/с)
	WindGust  float64 `json:"wind_gust"`  // Скорость порывов ветра (в м/с)
	WindDir   string  `json:"wind_dir"`   // Направление ветра. Возможные значения:
	//«nw» — северо-западное
	//«n» — северное
	//«ne» — северо-восточное
	//«e» — восточное
	//«se» — юго-восточное
	//«s» — южное
	//«sw» — юго-западное
	//«w» — западное
	//«с» — штиль
	PressureMm int    `json:"pressure_mm"` // Давление (в мм рт. ст.)
	PressurePa int    `json:"pressure_pa"` // Давление (в гектопаскалях)
	Humidity   int    `json:"humidity"`    // Влажность воздуха (в процентах)
	Daytime    string `json:"daytime"`     // Светлое или темное время суток. Возможные значения:
	//«d» — светлое время суток
	//«n» — темное время суток
	Polar  bool   `json:"polar"`  // Признак того, что время суток, указанное в поле daytime, является полярным
	Season string `json:"season"` // Время года в данном населенном пункте. Возможные значения:
	//«summer» — лето
	//«autumn» — осень
	//«winter» — зима
	//«spring» — весна
	ObsTime int64 `json:"obs_time"` // Время замера погодных данных в формате Unixtime
}

type Forecast struct {
	Date     string `json:"date"`      // Дата прогноза в формате ГГГГ-ММ-ДД
	DateTs   int64  `json:"date_ts"`   // Дата прогноза в формате Unixtime
	Week     int    `json:"week"`      // Порядковый номер недели
	Sunrise  string `json:"sunrise"`   // Время восхода Солнца, локальное время (может отсутствовать для полярных регионов)
	Sunset   string `json:"sunset"`    // Время заката Солнца, локальное время (может отсутствовать для полярных регионов)
	MoonCode int    `json:"moon_code"` // Код фазы Луны. Возможные значения:
	//0 — полнолуние
	//1-3 — убывающая Луна
	//4 — последняя четверть
	//5-7 — убывающая Луна
	//8 — новолуние
	//9-11 — растущая Луна
	//12 — первая четверть
	//13-15 — растущая Луна
	MoonText string `json:"moon_text"` // Текстовый код для фазы Луны. Возможные значения:
	//moon-code-0 — полнолуние
	//moon-code-1 — убывающая луна
	//moon-code-2 — убывающая луна
	//moon-code-3 — убывающая луна
	//moon-code-4 — последняя четверть
	//moon-code-5 — убывающая луна
	//moon-code-6 — убывающая луна
	//moon-code-7 — убывающая луна
	//moon-code-8 — новолуние
	//moon-code-9 — растущая луна
	//moon-code-10 — растущая луна
	//moon-code-11 — растущая луна
	//moon-code-12 — первая четверть
	//moon-code-13 — растущая луна
	//moon-code-14 — растущая луна
	//moon-code-15 — растущая луна
	Parts []Part `json:"parts"` // Прогнозы по времени суток
}

type Part struct {
	PartName string `json:"part_name"` // Название времени суток. Возможные значения:
	//night — ночь
	//morning — утро
	//day — день
	//evening — вечер
	TempMin   int    `json:"temp_min"`   // Минимальная температура для времени суток (°C)
	TempMax   int    `json:"temp_max"`   // Максимальная температура для времени суток (°C)
	TempAvg   int    `json:"temp_avg"`   // Средняя температура для времени суток (°C)
	FeelsLike int    `json:"feels_like"` // Ощущаемая температура (°C)
	Icon      string `json:"icon"`       // Код иконки погоды, иконка доступна по адресу https://yastatic.net/weather/i/icons/funky/dark/<значение из поля icon>.svg
	Condition string `json:"condition"`  // Код расшифровки погодного описания. Возможные значения:
	//clear — ясно
	//partly-cloudy — малооблачно
	//cloudy — облачно с прояснениями
	//overcast — пасмурно
	//drizzle — морось
	//light-rain — небольшой дождь
	//rain — дождь
	//moderate-rain — умеренно сильный дождь
	//heavy-rain — сильный дождь
	//continuous-heavy-rain — длительный сильный дождь
	//showers — ливень
	//wet-snow — дождь со снегом
	//light-snow — небольшой снег
	//snow — снег
	//snow-showers — снегопад
	//hail — град
	//thunderstorm — гроза
	//thunderstorm-with-rain — дождь с грозой
	//thunderstorm-with-hail — гроза с градом
	Daytime string `json:"daytime"` // Светлое или темное время суток. Возможные значения:
	//«d» — светлое время суток
	//«n» — темное время суток
	Polar     bool    `json:"polar"`      // Признак того, что время суток, указанное в поле daytime, является полярным
	WindSpeed float64 `json:"wind_speed"` // Скорость ветра (в м/с)
	WindGust  float64 `json:"wind_gust"`  // Скорость порывов ветра (в м/с)
	WindDir   string  `json:"wind_dir"`   // Направление ветра. Возможные значения:
	//«nw» — северо-западное
	//«n» — северное
	//«ne» — северо-восточное
	//«e» — восточное
	//«se» — юго-восточное
	//«s» — южное
	//«sw» — юго-западное
	//«w» — западное
	//«с» — штиль
	PressureMm int `json:"pressure_mm"` // Давление (в мм рт. ст.)
	PressurePa int `json:"pressure_pa"` // Давление (в гектопаскалях)
	Humidity   int `json:"humidity"`    // Влажность воздуха (в процентах)
	PrecMm     int `json:"prec_mm"`     // Прогнозируемое количество осадков (в мм)
	PrecPeriod int `json:"prec_period"` // Прогнозируемый период осадков (в минутах)
	PrecProb   int `json:"prec_prob"`   // Вероятность выпадения осадков
}
