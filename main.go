package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const notesFile = "notes.txt"

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: note <command>")
        fmt.Println("Commands: add, list, delete")
        return
    }

    command := os.Args[1]

    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Usage: note add <text>")
            return
        }
        text := strings.Join(os.Args[2:], " ")
        err := addNote(text)
        if err != nil {
            fmt.Println("Error: ", err)
        }
        fmt.Println("Note added successfully.")
    case "list":
        notes, err := listNotes()
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }
        fmt.Println("Notes:")
        for i, note := range notes {
            fmt.Printf("%d. %s\n", i+1, note)
        }
    case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Usage: note delete <note number>")
            return
        }
        noteNum := os.Args[2]
        err := deleteNote(noteNum)
        if err != nil {
            fmt.Println("Error: ", err)
        }
        fmt.Println("Note deleted successfully.")
    default:
        fmt.Println("Invalid command. Use 'add', 'list', or 'delete'.")
    }
}

func addNote(text string) error {
    file, err := os.OpenFile(notesFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    _, err = writer.WriteString(text + "\n")
    if err != nil {
        return err
    }

    err = writer.Flush()
    return err
}

func listNotes() ([]string, error) {
    file, err := os.Open(notesFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var notes []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        notes = append(notes, scanner.Text())
    }

    if scanner.Err() != nil {
        return nil, scanner.Err()
    }

    return notes, nil
}

func deleteNote(noteNum string) error {
    notes, err := listNotes()
    if err != nil {
        return err
    }

    n := len(notes)
    index := 0
    for i := 0; i < n; i++ {
        if noteNum == fmt.Sprintf("%d", i+1) {
            index = i
            break
        }
    }

    if index < n {
        notes = append(notes[:index], notes[index+1:]...)
        file, err := os.Create(notesFile)
        if err != nil {
            return err
        }
        defer file.Close()

        writer := bufio.NewWriter(file)
        for _, note := range notes {
            _, err := writer.WriteString(note + "\n")
            if err != nil {
                return err
            }
        }

        err = writer.Flush()
        return err
    } else {
        return fmt.Errorf("Note not found")
    }
}
