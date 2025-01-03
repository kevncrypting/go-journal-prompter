package main

import "fmt"
import "math/rand"

func main() {
	prompts := []string{
		"What is something that you're grateful for?",
		"What made you smile?",
		"Did you catch up with a friend today? What did you talk about?",
		"What did you do for yourself today?",
		"What do you want to focus on tomorrow?",
		"What is a challenge I faced recently, and how did I overcome it?",
		"How can I show myself self-care tomorrow?",
		"What emotions am I feeling right now?",
		"What qualities do I want to embody tomorrow?",
	}

	fmt.Println("Today's journal prompt:", prompts[rand.Intn(len(prompts))])
}
