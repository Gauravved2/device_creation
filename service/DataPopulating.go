package service

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	// "time"

	"bitbucket.org/klokinnovations/webapp/codegenerate"
	"bitbucket.org/klokinnovations/webapp/connection"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}

type returnStatus struct {
	Status  string
	Message string
}

func (s *Server) routes() {
	serv := s.PathPrefix("/api").Subrouter()
	serv.HandleFunc("/addDevice", s.AddDevice()).Methods("POST")
}

type AccessPointData struct {
	BuildingName    string `json:"Building"`
	Floor           string `json:"Floor"`
	Section         string `json:"Section"`
	AccessPointName string `json:"AccessPointName"`
	BSSID           string `gorm:"type:varchar(17);unique;not null"`
	SSID            string `gorm:"not null"`
}
type data struct {
	AccessPointData []AccessPointData
}

func (s *Server) AddDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAccessPointData AccessPointData
		var newData data
		if err := json.NewDecoder(r.Body).Decode(&newData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if len(newData.AccessPointData) == 0{
			http.Error(w, "Empty request body", http.StatusBadRequest)
		} else{
			for i, _ := range newData.AccessPointData {
				newAccessPointData = newData.AccessPointData[i]
				fmt.Println(newAccessPointData)
				
				dbConnection, err := connection.CreateConnection()
				if err != nil {
					http.Error(w, err.Error()+"Error in connecting to database", http.StatusInternalServerError)
				}
				buildingCode := codegenerate.GenerateCode(dbConnection,"buildings","B")
				floorCode := codegenerate.GenerateCode(dbConnection,"floors","F")
				sectionCode := codegenerate.GenerateCode(dbConnection,"sections","S")
				buildingShortName := getShortName(newAccessPointData.BuildingName)
				floorShortName := getShortName(newAccessPointData.Floor)
				sectionShortName := getShortName(newAccessPointData.Section)
				accessPointCode := codegenerate.GenerateCode(dbConnection,"access_points","AP")
				fmt.Println(buildingCode + " " + floorCode + " " + sectionCode + " " + buildingShortName)
				fmt.Println(floorShortName + " " + sectionShortName + accessPointCode)
				parsedBSSID, err := net.ParseMAC(newAccessPointData.BSSID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				fmt.Println("Connection Established")
				err = connection.AddBuilding(dbConnection, buildingCode, newAccessPointData.BuildingName, buildingShortName)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				err = connection.AddFloors(dbConnection, floorCode, buildingCode, newAccessPointData.Floor, floorShortName)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				err = connection.AddSections(dbConnection, sectionCode, newAccessPointData.Section, sectionShortName, floorCode)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				err = connection.AddAccessPoint(dbConnection, accessPointCode, newAccessPointData.AccessPointName, parsedBSSID.String(), newAccessPointData.SSID, floorCode, sectionCode)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				fmt.Println("Inserted")
				var returnStatusData returnStatus
				returnStatusData.Status = "OK"
				returnStatusData.Message = "Access Point Added"
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(returnStatusData); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				// rows, err := dbConnection.Query("SELECT code, name FROM buildings")
				// if err != nil {
				// 	log.Fatal(err)
				// }
				// defer rows.Close()
	
				// // Iterate over the rows
				// for rows.Next() {
				// 	var column1Value string
				// 	var column2Value string
	
				// 	// Scan the values from the current row into variables
				// 	err := rows.Scan(&column1Value, &column2Value)
				// 	if err != nil {
				// 		log.Fatal(err)
				// 	}
	
				// 	// Process the retrieved values
				// 	fmt.Println("Column1:", column1Value)
				// 	fmt.Println("Column2:", column2Value)
				// }
	
				// // Check for any errors during iteration
				// if err = rows.Err(); err != nil {
				// 	log.Fatal(err)
				// }
				defer dbConnection.Close()
			}
		}
	}
}

func getShortName(name string) string {
	shortName := ""
	part := strings.Split(name, " ")
	if len(part) == 1 {
		shortName = part[0]
	} else {
		for i := 0; i < len(part); i++ {
			n := part[i]
			if len(n) > 2 {
				shortName += n[:2]
			} else {
				shortName += n
			}
		}
	}
	return shortName
}
// func generateCode() string {
// 	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
// 	codeLength := 2
// 	code := make([]byte, codeLength)
// 	for i := 0; i < codeLength; i++ {
// 		// rand.Seed(time.Now().UnixNano())
// 		n := rand.Intn(len(charSet))
// 		// fmt.Println(n)
// 		code[i] = charSet[n]
// 	}
// 	return string(code)
// }
