/*
What it is: 
This is a set of packages to be marshaled from 'go' to .json. 
Why is: 
Standard unit of refusing to encode json from hash map keys which are not strings, 
This package encode them as arrays and hash map whose string keys as objects. 
Also realized rapid encoding json alternative github.com/pquerna/ffjson

features:   
- coding without error;    
- supported json.Marshaler interface.
- tag json supported

Что это: 
Это набор пакетов для маршалинга из go в .json.
Зачем(почему) это:
Стандартны модуль го отказывается кодировать в json хешмапы ключи которых не являются строками, 
этот пакет их кодирует как массивы, а хешмапы у которых ключи строки как объекты.
Также реализуется быстрое кодирование в json на основе генераторов - альтернатива github.com/pquerna/ffjson
*/