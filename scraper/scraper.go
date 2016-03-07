//~~~~~~~~~~~~~~~~~~~~~~//
// Make an HTTP request //
//~~~~~~~~~~~~~~~~~~~~~~//

resp, _ := http.Get(url)
bytes, _ := ioutil.ReadAll(resp.Body)

fmt.Println("HTML:\n\n", string(bytes))

resp.Body.Close()