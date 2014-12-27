package main

func Copy(newMap *map[string]string, oldMap map[string]string) {
	tmpMap := map[string]string{}

	for k, v := range oldMap {
		tmpMap[k] = v
	}

	*newMap = tmpMap
}

func main() {

}
