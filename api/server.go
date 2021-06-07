package main

import (
	// "bufio"
	// "fmt"
	// "io"
	"log"

	// "mime/multipart"
	// "os"
	// "strings"

	"github.com/gofiber/fiber/v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024 * 1024,
	})

	// app.Static("/results", "./results")

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// app.Post("/api", func(c *fiber.Ctx) error {
	// 	file, err := c.FormFile("document")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return c.SaveFile(file, fmt.Sprintf("./results/%s", file.Filename))
	// })

	// app.Post("/upload", func(c *fiber.Ctx) error {
	// 	// ctx.FormFile("fileName") ->
	// 	// multiHeader.File ->
	// 	// Open ->
	// 	// use the File object and ->
	// 	//  stream it in a new created file
	// 	file, err := c.FormFile("data")
	// 	if err != nil {
	// 		return err
	// 	}

	// 	fmt.Printf("%s", file.Filename)

	// 	infile_header, err := file.Open()
	// 	// infile_header := multipart.Open(file)

	// 	// make a read buffer
	// 	// r0 := multipart.NewReader(infile_header)

	// 	// make a read buffer
	// 	r := bufio.NewReader(infile_header)

	// 	// open output file
	// 	fo, err := os.Create(fmt.Sprintf("./results/%s", file.Filename))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// close fo on exit and check for its returned error
	// 	defer func() {
	// 		if err := fo.Close(); err != nil {
	// 			panic(err)
	// 		}
	// 	}()
	// 	// make a write buffer
	// 	w := bufio.NewWriter(fo)

	// 	// make a buffer to keep chunks that are read
	// 	buf := make([]byte, 1024)
	// 	for {
	// 		// read a chunk
	// 		n, err := r.Read(buf)
	// 		if err != nil && err != io.EOF {
	// 			panic(err)
	// 		}
	// 		if n == 0 {
	// 			break
	// 		}

	// 		// write a chunk
	// 		if _, err := w.Write(buf[:n]); err != nil {
	// 			panic(err)
	// 		}
	// 	}

	// 	if err = w.Flush(); err != nil {
	// 		panic(err)
	// 	}

	// 	return c.SendString(file.Filename)
	// })

	app.Post("/stream", func(c *fiber.Ctx) error {

		// b0 := c.Body()
		// fmt.Printf("b0: %s\nb1: %s\nb2: %s\n", b0, b1, b2)
		log.Printf("is request body stream? %v\n", c.Context().IsBodyStream())
		log.Printf("Request size: %f Go\n", float32(len(c.Body()))/1000000000.0)

		return c.SendString("end of server stream func.\n")
	})

	log.Fatal(app.Listen(":4000"))
}
