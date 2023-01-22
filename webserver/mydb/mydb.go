package mydb

//pgadmin password zander

import (
	"database/sql"
	"fmt"
	s "strings"
	"webserver/constants"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "Seaking"
)

var DB *sql.DB

func init() {
	OpenDB()
}

func OpenDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	DB = db
}

/*
func QueryDB(query string) (myrows []string) {
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Errorf("canot run query %s: %w", query, err)
		return
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	row := make([]interface{}, len(cols))
	rowPtr := make([]interface{}, len(cols))
	for i := range row {
		rowPtr[i] = &row[i]
	}
	fmt.Println(cols)
	b, _ := json.Marshal(cols)
	myrows = append(myrows, string(b))
	for rows.Next() {
		err = rows.Scan(rowPtr...)
		if err != nil {
			fmt.Println("cannot scan row:", err)
		}
		//fmt.Println(row...)
		b, _ := json.Marshal(row)

		//fmt.Println(string(b))
		myrows = append(myrows, string(b))
	}
	//return rows.Err()
	return
}
*/

func QueryData(name string, key string, limit string) (myrows []constants.SatData) {
	//query := "select id,mid,k,v,ts,t,name from satdata where name=' " + name + "' and key='" + key + "'  order by ts desc limit " + limit + ";"
	if len(limit) < 1 {
		limit = "5"
	}
	query := fmt.Sprintf(" select id,mid,key,val,ts,t,name from satdata where name = '%s' and key = '%s' order by ts desc limit %s", name, key, limit)
	fmt.Println(query)
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Errorf("cannot run query %s: %w", query, err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		satdata := constants.SatData{}
		err = rows.Scan(&satdata.ID, &satdata.MID, &satdata.Key, &satdata.Val, &satdata.TS, &satdata.T, &satdata.Name)

		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(satdata)

		myrows = append(myrows, satdata)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	//return rows.Err()
	return
}

func InsertData(Parameters map[string]string) { //Edits name to easily have as Key Value pair
	fmt.Print("Insert Data    ")
	fmt.Println(Parameters)

	mid := Parameters["UID"]
	if mid == "" {
		mid = Parameters["MID"]
	}
	t := Parameters["T"]

	name := Parameters["N"]
	//sqlStatement := `  INSERT INTO satdata (mid, k, v)  VALUES ($1, $2, $3)
	//id := 0
	//err := DB.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
	//err := DB.QueryRow(sqlStatement).Scan(&id)

	for k, v := range Parameters {
		if k != "mid" && k != "ts" {

			sqlStatement := fmt.Sprintf(" INSERT INTO satdata (mid, key, val, t, name)  VALUES ('%s','%s','%s','%s','%s' )", mid, k, v, t, name)
			fmt.Println(sqlStatement)

			_, err := DB.Exec(sqlStatement)
			if err != nil {
				fmt.Println("Error in InsertData")
			}
		}
	}
}

func InsertCMD(bsid, data, sendon string) {
	if len(sendon) < 1 {
		sendon = "NOW()"
	} else {
		sendon = "'" + sendon + "'"
	}
	//sqlStatement := `  INSERT INTO commands (bsid, sendon, data, needack)  VALUES ($1, $2, $3, $4) `
	//sqlStatement := `  INSERT INTO commands (bsid,  data, sendon)  VALUES ('` + bsid + `','` + data + `',` + sendon + `) `
	sqlStatement := fmt.Sprintf(" INSERT INTO commands (bsid,  data, sendon)  VALUES (now(),'%s','%s',%s)", bsid, data, sendon)
	fmt.Println(sqlStatement)

	//err := DB.QueryRow(sqlStatement, bsid, sendon, data, needack)
	_, err := DB.Exec(sqlStatement)
	if err != nil {
		fmt.Println("Error in InsertCmd")
	}
}

/*
func InsertImageName(Parameters map[string]string) {
	filename := Parameters["FILENAME"]
	totalblocks := Parameters["TOTALBLOCKS"]
	width := Parameters["WIDTH"]
	rssi := Parameters["RSSI"]
	snr := Parameters["SNR"]
	sqlStatement := fmt.Sprintf(" INSERT INTO satimages (ts, filename,  totalblocks, width, rssi, snr)  VALUES (now(),'%s','%s','%s', '%s', '%s')", filename, totalblocks, width, rssi, snr)
	fmt.Println(sqlStatement)

	//err := DB.QueryRow(sqlStatement, bsid, sendon, data, needack)
	_, err := DB.Exec(sqlStatement)
	if err != nil {
		fmt.Println("Error in InsertImageNAme")
	}
}
*/

//https://www.sanarias.com/blog/1214PlayingwithimagesinHTTPresponseingolang

func InsertImagePart(filename string, val []byte) {
	str := filename
	str = s.Replace(str, ".jpg", "", -1)
	//str = s.Replace(str, "img", "", -1)
	pos := s.Index(str, "_")

	part := str[len(str)-1:]
	filename = str[0:pos]
	block := str[pos+1 : len(str)-1]

	//need to split filename into components
	//sqlStatement := fmt.Sprintf(" INSERT INTO satimageparts (ts, filename, key, val)  VALUES (now(),'%s','%s',$1)", filename, filename)
	sqlStatement := fmt.Sprintf(" INSERT INTO satimages ( filename, block, part,data)  VALUES ('%s','%s','%s',$1)", filename, block, part)
	fmt.Println(sqlStatement)

	_, err := DB.Exec(sqlStatement, val)
	if err != nil {
		fmt.Println("Error in InsertImagePart")
		fmt.Println(err)
	}
}

func UpdateCMDAck(bsid, id string) {
	sqlStatement := `  update commands set ack=TRUE where id =  $1 `

	err := DB.QueryRow(sqlStatement, id)
	if err != nil {
		fmt.Println("Error in UpdateCMDAck")
	}
}

func UpdateCMDSent(bsid, id string) {
	sqlStatement := fmt.Sprintf(" update commands set sent=TRUE, sentby='%s', senton=now() where id =  %s", bsid, id)
	fmt.Println(sqlStatement)
	_, err := DB.Exec(sqlStatement)
	if err != nil {
		fmt.Println("Error in UpdateCMDSent")
	}
}

func Register() (myrows []constants.Stations) {
	var id string
	sqlStatement := `  insert into stations(name) values('') returning id `

	err := DB.QueryRow(sqlStatement).Scan(&id)

	if err != nil {
		fmt.Println("Error in Register")
	}

	stations := constants.Stations{}
	stations.ID = id

	myrows = append(myrows, stations)
	return
}

func GetImage(filename, block string) (image []byte) {
	//encode(coalesce(part0, '')||coalesce(part1, '')||coalesce(part2, ''), 'escape')as image    //This is turn it into txt to see in sql
	sqlStatement := fmt.Sprintf(" select filename, block, part, data from satimages where filename='%s' and block='%s' and block!='_' order by  part asc limit 4", filename, block)
	fmt.Println(sqlStatement)
	rows, err := DB.Query(sqlStatement)
	if err != nil {
		fmt.Errorf("cannot run query %s: %w", sqlStatement, err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		satimg := constants.SatImages{}
		//err = rows.Scan(&satimg.Id, &satimg.Filename, &satimg.Block, &satimg.Parts, &satimg.Totalblocks, &satimg.Width, &satimg.Ts, &satimg.Part0, &satimg.Part1, &satimg.Part2, &satimg.Part3, &satimg.Part4)
		err = rows.Scan(&satimg.Filename, &satimg.Block, &satimg.Part, &satimg.Data)

		if err != nil {
			fmt.Println("Error in GetImage")
			panic(err)
		}
		//fmt.Println(satimg)

		image = append(image, satimg.Data...)

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Errorf("cannot run getImage")
	}

	//return rows.Err()
	return
}

func GetCMD(bsid string) (mycmds []constants.Commands) {
	query := "select text(id),bsid,sendon,replace(replace(data,':',';'),'~','^'),text(sent),text(ack),text(needack), ts from commands  where (sent=false ) and (sendon<now()) and (bsid='all' or bsid='" + bsid + "') order by sendon desc limit 1;"
	fmt.Println(query)
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Errorf("cannot run query %s: %w", query, err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		satcmd := constants.Commands{}
		err = rows.Scan(&satcmd.Id, &satcmd.BSid, &satcmd.Sendon, &satcmd.Data, &satcmd.Sent, &satcmd.Ack, &satcmd.NeedAck, &satcmd.Ts)

		if err != nil {
			fmt.Println("Error in GetCMD")
			panic(err)
		}
		fmt.Println(satcmd)

		mycmds = append(mycmds, satcmd)
		UpdateCMDSent(bsid, satcmd.Id)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	//return rows.Err()
	return
}

func GetGPS(bsid string) (mycmds []constants.GPSData) {
	//query := "select text(id),bsid,sendon,replace(replace(data,':',';'),'~','^'),text(sent),text(ack),text(needack), ts from commands  where (sent=false ) and (sendon<now()) and (bsid='all' or bsid='" + bsid + "') order by sendon desc limit 1;"
	query := "SELECT e.val as lat, m.val as lon,e.ts  FROM satdata e INNER JOIN satdata m ON m.mid = e.mid and e.key='LAT' and m.key='LON' and length(m.mid)>0 and m.mid in( 	select mid from satdata where val='" + bsid + "') order by e.ts asc"

	fmt.Println(query)
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Errorf("cannot run query %s: %w", query, err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		gpsdata := constants.GPSData{}
		err = rows.Scan(&gpsdata.Lat, &gpsdata.Lon, &gpsdata.TS)

		if err != nil {
			fmt.Println("Error in GetCMD")
			panic(err)
		}
		fmt.Println(gpsdata)

		mycmds = append(mycmds, gpsdata)

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	//return rows.Err()
	return
}

func GetMSG(mid string) (Parameters map[string]string) {
	//query := "select text(id),bsid,sendon,replace(replace(data,':',';'),'~','^'),text(sent),text(ack),text(needack), ts from commands  where (sent=false ) and (sendon<now()) and (bsid='all' or bsid='" + bsid + "') order by sendon desc limit 1;"
	query := fmt.Sprintf(" select id,mid,key,val,ts,t,name from satdata where mid = '%s' ", mid)
	fmt.Println(query)
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Errorf("cannot run query %s: %w", query, err)
		return
	}

	Parameters = make(map[string]string)

	defer rows.Close()
	for rows.Next() {
		satdata := constants.SatData{}
		err = rows.Scan(&satdata.ID, &satdata.MID, &satdata.Key, &satdata.Val, &satdata.TS, &satdata.T, &satdata.Name)

		if err != nil {
			// handle this error
			panic(err)
		}

		if len(satdata.Key) > 0 {
			k := satdata.Key
			v := satdata.Val
			Parameters[k] = v
			fmt.Println(k)
			fmt.Println(v)
		}
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return Parameters
}
