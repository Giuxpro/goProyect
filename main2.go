package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func list_all(folder_name) {
	return
}

func parse_data(data_lines) {
	data  := `{
        "To": null,
        "From": null,
        "Subject": null,
        "Body": ""
    }`

	for _, line := range data_lines {
		if line.find("Message-ID:", 0) == 0{
            data[Message-ID] = line[9:len(line)]
		}else if line.find("Date:", 0) == 0{
            data[Date] = line[9:len(line)]
		}else if line.find("From:", 0) == 0{
            data[From] = line[6:len(line)]
		}else if line.find("To:", 0) == 0{
            data[To] = line[4:len(line)]
        }else if line.find("Subject:", 0) == 0{
            data[Subject] = line[9:len(line)]
        }else if line.find("Cc:", 0) == 0{
            data[Cc] = line[9:len(line)]
        }else if line.find("Mime-Version:", 0) == 0{
            data[Mime-Version] = line[9:len(line)]
        }else if line.find("Content-Type:", 0) == 0{
            data[Content-Type] = line[9:len(line)]
        }else if line.find("Content-Transfer-Encoding:", 0) == 0{
            data[Content-Transfer-Encoding] = line[9:len(line)]
        }else if line.find("X-From:", 0) == 0{
            data[X-From] = line[9:len(line)]
        }else if line.find("X-To:", 0) == 0{
            data[X-To] = line[9:len(line)]
        }else if line.find("X-cc:", 0) == 0{
            data[X-cc] = line[9:len(line)]
        }else if line.find("X-bcc:", 0) == 0{
            data[X-bcc] = line[9:len(line)]
        }else if line.find("X-Folder:", 0) == 0{
            data[X-Folder] = line[9:len(line)]
        }else if line.find("X-Origin:", 0) == 0{
            data[X-Origin] = line[9:len(line)]
        }else if line.find("X-FileName:", 0) == 0{
            data[X-FileName] = line[9:len(line)]
        }else{
            data[Body] = data[Body] + line
		}
    return data
		
	}
}


func index_data(data){

	req, err := http.NewRequest("POST", "http://localhost:4080/api/games3/_doc", strings.NewReader(data)) ///api/games3/_doc
    fmt.Print(data)
    if err != nil {
        log.Fatal(err)
    }
    req.SetBasicAuth("admin", "Complexpass#123")
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    log.Println(resp.StatusCode)
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(body))
}

func main() {
	maildir := "./enron_mail_20110402/maildir"
	user_list := list_all(maildir)

	for _, user := range user_list {
		folders := list_all(maildir + "/" + user)
		for _, folder := range folders {
			mail_files = list_files(maildir + user + "/" + folder + "/")
			for _, mail_file := range mail_files {
				fmt.Println("Indexing : ", user + "/" + folder + "/" + mail_file)
				sys_file = open(maildir + "/" + user + "/" + folder + "/" + mail_file, "r")
				lines = sys_file.ReadAll()

				index_data(parse_data(lines))
			}
		}
	}
}
