package myweb

import (
	"encoding/json"
	"fmt"
	"image/png"
	"io"
	"net/http"
	"os"
	"time"
	"webserver/mydb"
	ir "webserver/myimagemaker"
)

func pageTime(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func pageQuery(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	var name, key, limit string
	name = values.Get("NAME")
	key = values.Get("KEY")
	limit = values.Get("LIMIT")

	data := mydb.QueryData(name, key, limit)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func pageInsert(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("get the form data")
	fmt.Println(r.PostForm)

	Parameters := make(map[string]string)

	for key, value := range r.PostForm {
		if len(value[0]) > 0 && len(key) > 0 {
			Parameters[key] = value[0]
			fmt.Println(key, " => ", value[0])
		}
	}

	mydb.InsertData(Parameters)
	/*
		values := r.URL.Query()
		var data, rssi, snr string

		data = r.FormValue("DATA")
		rssi = r.FormValue("RSSI")
		snr = r.FormValue("SNR")
		//data = values.Get("DATA")
		//rssi = values.Get("RSSI")
		//snr = values.Get("SNR")

		fmt.Println("pageInsert")
		mydb.LogDB(data, rssi, snr)
		for k, v := range values {
			fmt.Println(k, " => ", v)
			w.Write([]byte(k + " => " + v[0] + "\n"))
			w.Write([]byte(v[0] + "\n"))
			//w.Write([]byte(v[1] + "\n"))

			Parameters := myserial.Deserialize(v[0])
			fmt.Println(myserial.Serialize(Parameters))
			w.Write([]byte(myjson.CreateStrJSON(Parameters)))

			mydb.InsertData(Parameters)
		}
	*/
}

func pageInsertMulti(w http.ResponseWriter, r *http.Request) {
	uploadPath := "./upload"

	fmt.Println(".................PageInsertMulti.........................")
	err := r.ParseMultipartForm(10)
	if err != nil {
		fmt.Println("Parse ERROR")
		return
	}
	mForm := r.MultipartForm

	for k, _ := range mForm.File {

		fmt.Println(k)
		// k is the key of file part
		file, fileHeader, err := r.FormFile(k)

		if err != nil {
			fmt.Println("invoke FormFile error:", err)
			return
		}
		defer file.Close()

		fmt.Printf("the uploaded file: name[%s], size[%d], header[%#v]\n", fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		// store uploaded file into local path
		localFileName := uploadPath + "/" + fileHeader.Filename
		out, err := os.Create(localFileName)

		if err != nil {
			fmt.Printf("failed to open the file %s for writing", localFileName)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)

		if err != nil {
			fmt.Printf("copy file err:%s\n", err)
			return
		}

		fmt.Printf("file %s uploaded ok\n", fileHeader.Filename)

		var buffer []byte
		buffer = make([]byte, fileHeader.Size)
		file.Seek(0, 0)
		file.Read(buffer)

		for i := 0; i < len(buffer); i++ {
			fmt.Printf("%c", buffer[i])
		}
		mydb.InsertImagePart(fileHeader.Filename, buffer)
	}
	fmt.Println("pageInsertMulti")
}

/*
func pageInsertImageName(w http.ResponseWriter, r *http.Request) {

		values := r.URL.Query()
		//	var data, rssi, snr string
		//data = values.Get("DATA")
		//rssi = values.Get("RSSI")
		//snr = values.Get("SNR")

		fmt.Println("pageInsert")
		for k, v := range values {
			fmt.Println(k, " => ", v)
			w.Write([]byte(k + " => " + v[0] + "\n"))
			w.Write([]byte(v[0] + "\n"))
			//w.Write([]byte(v[1] + "\n"))

			Parameters := myserial.Deserialize(v[0])
			fmt.Println(myserial.Serialize(Parameters))
			w.Write([]byte(myjson.CreateStrJSON(Parameters)))

			mydb.InsertImageName(Parameters)
		}
	}
*/
func pageInsertCmd(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	var bsid, sendon, data string

	bsid = values.Get("BSID")
	data = values.Get("DATA")
	sendon = values.Get("SENDON")

	mydb.InsertCMD(bsid, data, sendon)
}

func pageGetCmd(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	bsid := values.Get("BSID")

	data := mydb.GetCMD(bsid)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func pageGetGPS(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	bsid := values.Get("BSID")

	data := mydb.GetGPS(bsid)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func pageGetMSG(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	mid := values.Get("MID")

	data := mydb.GetMSG(mid)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func pageGetImage(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	filename := values.Get("filename")
	block := values.Get("block")

	data := mydb.GetImage(filename, block)
	fmt.Println(filename)
	fmt.Println(block)
	fmt.Println(len(data))

	for i := 0; i < len(data); i++ {
		fmt.Print(data[i])
	}

	w.Header().Set("Content-Type", "image/jpeg")
	//w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	//w.Header().Set("Content-Disposition", `attachment;filename="sid.jpg"`)
	if _, err := w.Write(data); err != nil {
		fmt.Println("unable to write image.")
	}
}

func pageGetIRImage(w http.ResponseWriter, r *http.Request) {

	//new_png_file := "./upload/chessboard.png" // output image lives here
	values := r.URL.Query()
	filename := values.Get("filename")

	data := mydb.GetLastImage(filename)

	fmt.Println(filename)

	fmt.Println(len(data))

	for i := 0; i < len(data); i++ {
		fmt.Print(data[i])
	}

	if len(data) > 0 {
		myimage := ir.MakeIRImage(data)
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, myimage)
	}

	/*
		myfile, err := os.Create(new_png_file)
		if err != nil {
			panic(err.Error())
		}
		defer myfile.Close()
		png.Encode(myfile, myimage)                 // ... save image
		fmt.Println("new ir immage ", new_png_file) // view image issue : firefox  /tmp/chessboard.png
	*/

}

func pageRegister(w http.ResponseWriter, r *http.Request) {
	//values := r.URL.Query()
	//bsid := values.Get("BSID")
	//data := mydb.GetCMD(bsid)
	data := mydb.Register()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func pageUpdateCmd(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	var bsid, cid, ack, sent string

	bsid = values.Get("BSID")
	cid = values.Get("CID")
	ack = values.Get("ACK")
	sent = values.Get("SENT")

	if ack == "1" {
		mydb.UpdateCMDAck(bsid, cid)
	}

	if sent == "1" {
		mydb.UpdateCMDSent(bsid, cid)
	}
	w.Write([]byte("pageUpdateCmd"))
}

func ListenWeb() {
	mux := http.NewServeMux()
	// Convert the timeHandler function to a http.HandlerFunc type.

	// And add it to the ServeMux.
	mux.Handle("/time", http.HandlerFunc(pageTime))
	mux.Handle("/register", http.HandlerFunc(pageRegister))       //
	mux.Handle("/insert", http.HandlerFunc(pageInsert))           //
	mux.Handle("/insertmulti", http.HandlerFunc(pageInsertMulti)) //
	mux.Handle("/query", http.HandlerFunc(pageQuery))             //
	mux.Handle("/insertcmd", http.HandlerFunc(pageInsertCmd))     //
	mux.Handle("/getcmd", http.HandlerFunc(pageGetCmd))           //
	mux.Handle("/getgps", http.HandlerFunc(pageGetGPS))
	mux.Handle("/getmsg", http.HandlerFunc(pageGetMSG))
	mux.Handle("/updatecmd", http.HandlerFunc(pageUpdateCmd)) //

	mux.Handle("/image", http.HandlerFunc(pageGetImage))     //
	mux.Handle("/irimage", http.HandlerFunc(pageGetIRImage)) //

	http.ListenAndServe(":80", mux)
}

//http://192.168.86.148/image?filename=img17&block=229
//http://192.168.86.148/irimage?filename=IRZ1
//http://192.168.86.148/query?NAME=SATINFO&KEY=TEMPOBC

//http://192.168.86.148/insertcmd?BSID=&DATA=&SENDON
