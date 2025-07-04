package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/duixe/social_app/internal/models"
	"github.com/duixe/social_app/internal/repository"
)

var firstnames = []string{
	"John", "Alice", "Michael", "Sophia", "James", "Emma", "Oliver", "Isabella", "Liam", "Mia",
	"William", "Charlotte", "Ethan", "Amelia", "Mason", "Harper", "Noah", "Evelyn", "Lucas", "Abigail",
	"Henry", "Ella", "Benjamin", "Avery", "Alexander", "Scarlett", "Daniel", "Madison", "Sebastian", "Aria",
	"Matthew", "Luna", "Jackson", "Chloe", "Logan", "Grace", "David", "Ellie", "Samuel", "Hannah",
	"Joseph", "Zoey", "Carter", "Riley", "Owen", "Lily", "Elijah", "Layla", "Aiden", "Nora",
}

var lastnames = []string{
	"Smith", "Johnson", "Brown", "Williams", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
	"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
	"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson",
	"Walker", "Young", "Allen", "King", "Scott", "Green", "Adams", "Baker", "Nelson", "Hill",
	"Hall", "Rivera", "Campbell", "Mitchell", "Carter", "Roberts", "Gomez", "Phillips", "Evans", "Turner",
}

var titles = []string{
	"The Power of Habit", "Embracing Minimalism", "Healthy Eating Tips",
	"Travel on a Budget", "Mindfulness Meditation", "Boost Your Productivity",
	"Home Office Setup", "Digital Detox", "Gardening Basics",
	"DIY Home Projects", "Yoga for Beginners", "Sustainable Living",
	"Mastering Time Management", "Exploring Nature", "Simple Cooking Recipes",
	"Fitness at Home", "Personal Finance Tips", "Creative Writing",
	"Mental Health Awareness", "Learning New Skills",
}

var contents = []string{
	"In this post, we'll explore how to develop good habits that stick and transform your life.",
	"Discover the benefits of a minimalist lifestyle and how to declutter your home and mind.",
	"Learn practical tips for eating healthy on a budget without sacrificing flavor.",
	"Traveling doesn't have to be expensive. Here are some tips for seeing the world on a budget.",
	"Mindfulness meditation can reduce stress and improve your mental well-being. Here's how to get started.",
	"Increase your productivity with these simple and effective strategies.",
	"Set up the perfect home office to boost your work-from-home efficiency and comfort.",
	"A digital detox can help you reconnect with the real world and improve your mental health.",
	"Start your gardening journey with these basic tips for beginners.",
	"Transform your home with these fun and easy DIY projects.",
	"Yoga is a great way to stay fit and flexible. Here are some beginner-friendly poses to try.",
	"Sustainable living is good for you and the planet. Learn how to make eco-friendly choices.",
	"Master time management with these tips and get more done in less time.",
	"Nature has so much to offer. Discover the benefits of spending time outdoors.",
	"Whip up delicious meals with these simple and quick cooking recipes.",
	"Stay fit without leaving home with these effective at-home workout routines.",
	"Take control of your finances with these practical personal finance tips.",
	"Unleash your creativity with these inspiring writing prompts and exercises.",
	"Mental health is just as important as physical health. Learn how to take care of your mind.",
	"Learning new skills can be fun and rewarding. Here are some ideas to get you started.",
}

var tags = []string{
	"Self Improvement", "Minimalism", "Health", "Travel", "Mindfulness",
	"Productivity", "Home Office", "Digital Detox", "Gardening", "DIY",
	"Yoga", "Sustainability", "Time Management", "Nature", "Cooking",
	"Fitness", "Personal Finance", "Writing", "Mental Health", "Learning",
}

var comments = []string{
	"Great post! Thanks for sharing.",
	"I completely agree with your thoughts.",
	"Thanks for the tips, very helpful.",
	"Interesting perspective, I hadn't considered that.",
	"Thanks for sharing your experience.",
	"Well written, I enjoyed reading this.",
	"This is very insightful, thanks for posting.",
	"Great advice, I'll definitely try that.",
	"I love this, very inspirational.",
	"Thanks for the information, very useful.",
}

func Seed(repository repository.Repository, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)
	for _, user := range users {
		if err := repository.Users.Create(ctx, tx, user); err != nil {
			log.Println("Error creating user", err)
			return
		}
	}
	
	posts := generatePosts(200, users);
	for _, post := range posts {
		if err := repository.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post ", err )
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := repository.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment ", err )
			return
		}
	}

	log.Println("Seeding completed")
	
}

func generateUsers(num int) []*models.User {
	users := make([]*models.User, num)

	for i := 0; i < num; i++ {
		firstName := firstnames[i%len(firstnames)]
		lastName := lastnames[i%len(lastnames)]

		username := strings.ToLower(fmt.Sprintf("%s_%s", firstName, lastName))
		email := fmt.Sprintf("%s.%sa%d@example.com", strings.ToLower(firstName), strings.ToLower(lastName), i)

		users[i] = &models.User{
			FirstName: firstName,
			LastName:  lastName,
			Username:  username,
			Email:     email,
		}

		users[i].Password.Set("password")
	}

	return users
}

func generatePosts(num int, users []*models.User) []*models.Post {
	posts := make([]*models.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &models.Post{
			UserId: user.ID,
			Title: titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				titles[rand.Intn(len(titles))],
				titles[rand.Intn(len(titles))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*models.User, posts []*models.Post) []*models.Comment {
	cms := make([]*models.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &models.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
