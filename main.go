package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "multi-lane-broadcast"
	broker := "localhost:9092"
	partitionCount := 3

	// 1. PRODUCER: Send 5 fruits
	fmt.Println("--- PRODUCER: Sending fruits ---")
	writer := &kafka.Writer{
		Addr:  kafka.TCP(broker),
		Topic: topic,
	}

	fruits := []string{
		"Açaí",
		"Ackee",
		"African cucumber", // Kiwano
		"Ambarella",
		"Apple",
		"Apricot",
		"Atemoya",
		"Australian finger lime",
		"Avocado",
		"Babaco",
		"Bael",
		"Banana",
		"Baobab",
		"Barbadine",
		"Bilberry",
		"Biriba",
		"Black sapote",
		"Blackberry",
		"Blueberry",
		"Boysenberry",
		"Breadfruit",
		"Cacao",
		"Caimito",
		"Calamansi",
		"Camu camu",
		"Canistel",
		"Cantaloupe",
		"Carambola", // Starfruit
		"Cashew apple",
		"Cherimoya",
		"Cherry",
		"Chico",
		"Chilean guava",
		"Chinese jujube",
		"Clementine",
		"Cloudberry",
		"Coconut",
		"Cranberry",
		"Cupuaçu",
		"Currant",
		"Date",
		"Davidson's plum",
		"Dragonfruit",
		"Durian",
		"Elderberry",
		"Feijoa",
		"Fig",
		"Gac",
		"Genip",
		"Gooseberry",
		"Governor's plum",
		"Grape",
		"Grapefruit",
		"Grumichama",
		"Guarana",
		"Guava",
		"Hog plum",
		"Honeydew",
		"Huckleberry",
		"Iceberg", // possibly a melon variety
		"Imbe",
		"Indian gooseberry",
		"Jaboticaba",
		"Jackfruit",
		"Jambul",
		"Japanese plum",
		"Jujube",
		"Kakadu plum",
		"Kiwi",
		"Kiwano", // horned melon
		"Korlan",
		"Kumquat",
		"Langsat",
		"Lemon",
		"Lime",
		"Lingonberry",
		"Loguat",
		"Longan",
		"Lucuma",
		"Lychee",
		"Mabolo",
		"Mamey sapote",
		"Mamoncillo",
		"Mango",
		"Mangosteen",
		"Marang",
		"Marula",
		"Medlar",
		"Melinjo",
		"Monkey fruit",
		"Mountain apple",
		"Mulberry",
		"Murici",
		"Nance",
		"Naranjilla",
		"Natal plum",
		"Nectarine",
		"Noni",
		"Olive",
		"Orange",
		"Oregon grape",
		"Otaheite gooseberry",
		"Papaya",
		"Passionfruit",
		"Pawpaw",
		"Peach",
		"Pear",
		"Pepino",
		"Pequi",
		"Persimmon",
		"Peruvian groundcherry",
		"Phalsa",
		"Physalis",
		"Pineapple",
		"Pineberry",
		"Pitanga",
		"Platonia",
		"Plum",
		"Plumcot",
		"Poha",
		"Pomegranate",
		"Pomelo",
		"Prickly pear",
		"Quince",
		"Rambai",
		"Rambutan",
		"Rangpur",
		"Raspberry",
		"Red banana",
		"Rose apple",
		"Rose hip",
		"Safou",
		"Salak",
		"Salal berry",
		"Santol",
		"Sapodilla",
		"Sapote",
		"Sea buckthorn",
		"Serviceberry",
		"Soursop",
		"Spanish lime",
		"Star apple",
		"Starfruit",
		"Strawberry",
		"Strawberry guava",
		"Sugar apple",
		"Sunberry",
		"Tamarillo",
		"Tamarind",
		"Tangerine",
		"Tangor",
		"Thimbleberry",
		"Tomato",
		"Ugli fruit",
		"Velvet apple",
		"Voavanga",
		"Water apple",
		"Watermelon",
		"Wax apple",
		"White sapote",
		"Wolfberry",
		"Wood apple",
		"Yangmei",
		"Yellow mombin",
		"Youngberry",
	}
	for i, f := range fruits {
		targetPartition := i % partitionCount
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Partition: targetPartition,
				Value:     []byte(f),
			},
		)
		if err != nil {
			log.Printf("Failed to send %s: %v", f, err)
		} else {
			fmt.Printf("Sent %s to Partition %d\n", f, targetPartition)
		}
	}

	writer.Close()
	fmt.Println("--- All fruits sent successfully ---")

}
