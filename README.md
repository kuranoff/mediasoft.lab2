Программа для расчета меры семантической близости векторов (Cosine Similarity)

1. Пользователь загружает датасет
2. Выполняется текстовый препроцессинг датасета (удаляются избыточные символы)
3. fastText формирует модель (bin) и словарь векторного представления (vec)
4. Программа парсит словарь (vec): извлекает слово и его векторное представление. Результат записывается в БД
5. Пользователь (с главной страницы) вводит два слова (из датасета) и рассчитывается их семантическая близость. Программа извлекает слова и вектора из БД. Преобразует массив векторов из типа string в float64. Выполняет операции над векторами.

См. директорию screenshots.

В качестве БД используется MongoDB.
В файле mongodb.model описана структура документов коллекций.  

Внешние зависимости:
- [Gonum](https://github.com/gonum/gonum) gonum.org/v1/gonum/mat
- [Prometheus Go client library](https://github.com/prometheus/client_golang) github.com/prometheus/client_golang/prometheus/promhttp
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) go.mongodb.org/mongo-driver/mongo
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) go.mongodb.org/mongo-driver/mongo/options
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) go.mongodb.org/mongo-driver/bson
- [Google UUID](https://github.com/google/uuid) github.com/google/uuid
- [Go Text](https://github.com/golang/text) golang.org/x/text/transform
- [Go Text](https://github.com/golang/text) golang.org/x/text/unicode/norm

Дополнительные зависимости:
[fastText](https://fasttext.cc/)
