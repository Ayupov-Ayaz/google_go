package web

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
)
/**
 	Upload files
 */
func UploadFileStart() {
	http.HandleFunc("/", showUploadFormHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8000", nil)

}

func showUploadFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<html>
			<body>
				<form action="/upload" method="post" enctype="multipart/form-data">
				    <label for="my_file">File:
				        <input id="my_file" type="file" name="my_file">
				    </label>
				
				    <input type="submit" value="Загрузить">
				</form>
			</body>
		</html>
	`))
}


func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// парсим форму, всё что больше 5мб будет положено во временную директорию
	r.ParseMultipartForm(5 * 1024 * 1024)
	// получаю файл и fileHeader
	file, fileHeader, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close() // Закрываем файл, что бы не текли ресурсы
	fmt.Println("Header.FileName: ", fileHeader.Filename)
	fmt.Println("Header.Header: ", fileHeader.Header)

	hasher := md5.New()
	// хеширование
	io.Copy(hasher, file)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(hasher.Sum(nil))

	}
