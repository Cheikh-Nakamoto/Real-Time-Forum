package implementations

import (
	"bim/internal/metier/models"
	"bim/internal/pkg"

	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PostService struct {
	db *sql.DB
}

func (p *PostService) GetDB() *sql.DB {
	return p.db
}
func (p *PostService) SetDB(db *sql.DB) {
	p.db = db
}
func (p *PostService) CreatePost(title, content, image string, userId uuid.UUID) error {
	post := models.Post{}
	post.SetTitle(title)
	post.SetContent(content)
	post.SetImage(image)
	post.SetUserId(userId)
	_, err := p.db.Exec(`INSERT INTO posts (title, content, image, user_id) VALUES (?, ?, ?, ?)`, post.GetTitle(), post.GetContent(), post.GetImage(), post.GetUserId().String())
	if err != nil {
		fmt.Println("erreur lors de la création du post dans la base de données")
		return err
	}
	return nil
}
func (p *PostService) UpdatePost(id uint, post models.Post, title, content, image string) error {
	post.SetTitle(title)
	post.SetContent(content)
	post.SetImage(image)
	post.SetUpdatedAt(time.Now())
	_, err := p.db.Exec(`UPDATE posts SET title = ?, content = ?, image = ?, updated_at = ? WHERE id = ?`, post.GetTitle(), post.GetContent(), post.GetImage(), post.GetUpdatedAt(), id)
	if err != nil {
		return err
	}
	return nil
}
func (p *PostService) DeletePost(post models.Post) error {
	_, err := p.db.Exec("DELETE FROM posts WHERE id = ?", post.GetID())
	if err != nil {
		return err
	}
	return nil
}
func (p *PostService) GetAllPosts() ([]models.Post, error) {
	rows, err := p.db.Query(`SELECT * FROM posts`)
	if err != nil {
		fmt.Println("aucun resultat pour la requête", err)
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var (
			post        models.Post
			id          uint
			title       string
			content     string
			image       string
			userId      string
			createdAt   time.Time
			userService UserService
			// categoryService CategoryService
			// commentService  CommentService
			// actionService   ActionService
		)
		userService.SetDB(p.GetDB())
		// categoryService.SetDB(p.GetDB())
		// commentService.SetDB(p.GetDB())
		// actionService.SetDB(p.GetDB())
		err = rows.Scan(&id, &title, &content, &image, &userId, &createdAt, &userId)

		if err != nil {
			fmt.Println("Erreur lors de la recuperation des donnes", err)
			return nil, err
		}

		userID, err := pkg.StringToUUID(userId)
		if err != nil {
			return nil, err
		}
		//fmt.Println(userID)
		// owner, err := userService.GetUserById(userID)
		// if err != nil {
		// 	return nil, err
		// }
		// fmt.Println(owner)
		// cats, err := categoryService.GetCategoriesByPost(id)
		// if err != nil {
		// 	return nil, err
		// }
		// comments, _ := commentService.GetCommentsIdByPosts(id)
		// comment, _ := commentService.CountCommentByPost(id)
		// likes, _ := actionService.GetLikesByPost(id)
		// dislikes, _ := actionService.GetDislikesByPost(id)

		post.SetID(id)
		post.SetTitle(title)
		post.SetContent(content)
		post.SetImage(image)
		post.SetUserId(userID)
		post.SetCreatedAt(createdAt)
		// post.SetOwner(owner.GetName())
		//  post.SetCategories(cats)
		// post.SetCommentsId(comments)
		// post.SetLike(likes)
		// post.SetDislike(dislikes)
		// post.SetComment(comment)
		posts = append(posts, post)
	}
	return posts, nil
}
func (p *PostService) GetPost(id uint) (models.Post, error) {
	var (
		post      models.Post
		ID        uint
		title     string
		content   string
		image     string
		userId    uuid.UUID
		createdAt time.Time
		updatedAt time.Time
		// userService     UserService
		// categoryService CategoryService
		// commentService  CommentService
		// actionService   ActionService
	)
	// userService.SetDB(p.GetDB())
	// categoryService.SetDB(p.GetDB())
	// commentService.SetDB(p.GetDB())
	// actionService.SetDB(p.GetDB())
	err := p.db.QueryRow(`SELECT id, title, content, image, user_id, created_at, updated_at FROM posts WHERE id=?`, id).Scan(&ID, &title, &content, &image, &userId, &createdAt, &updatedAt)
	if err != nil {
		return models.Post{}, err
	}
	// owner, err := userService.GetUserById(userId)
	// if err != nil {
	// 	return post, err
	// }
	// cats, err := categoryService.GetCategoriesByPost(id)
	// if err != nil {
	// 	return post, err
	// }
	// comments, _ := commentService.GetCommentsIdByPosts(id)
	// likes, _ := actionService.GetLikesByPost(id)
	// dislikes, _ := actionService.GetDislikesByPost(id)
	// comment, _ := commentService.CountCommentByPost(ID)
	post.SetID(ID)
	post.SetTitle(title)
	post.SetContent(content)
	post.SetImage(image)
	post.SetUserId(userId)
	post.SetCreatedAt(createdAt)
	post.SetUpdatedAt(updatedAt)
	// post.SetOwner(owner.GetName())
	// post.SetCategories(cats)
	// post.SetCommentsId(comments)
	// post.SetLike(likes)
	// post.SetDislike(dislikes)
	// post.SetComment(comment)
	return post, nil
}

//	func (p *PostService) GetByUser(userId uuid.UUID) ([]models.Post, error) {
//		rows, err := p.db.Query(`SELECT id, title, content, image, user_id, created_at, updated_at FROM posts WHERE user_id=?`, userId.String())
//		if err != nil {
//			return nil, err
//		}
//		defer rows.Close()
//		var posts []models.Post
//		for rows.Next() {
//			var (
//				post            models.Post
//				ID              uint
//				title           string
//				content         string
//				image           string
//				userId          uuid.UUID
//				createdAt       time.Time
//				updatedAt       time.Time
//				userService     UserService
//				categoryService CategoryService
//				commentService  CommentService
//				actionService   ActionService
//			)
//			userService.SetDB(p.GetDB())
//			categoryService.SetDB(p.GetDB())
//			commentService.SetDB(p.GetDB())
//			actionService.SetDB(p.GetDB())
//			err = rows.Scan(&ID, &title, &content, &image, &userId, &createdAt, &updatedAt)
//			if err != nil {
//				return nil, err
//			}
//			owner, err := userService.GetUserById(userId)
//			if err != nil {
//				return nil, err
//			}
//			cats, err := categoryService.GetCategoriesByPost(ID)
//			if err != nil {
//				return nil, err
//			}
//			comments, _ := commentService.GetCommentsByPosts(ID)
//			likes, _ := actionService.GetLikesByPost(ID)
//			dislikes, _ := actionService.GetDislikesByPost(ID)
//			comment, _ := commentService.CountCommentByPost(ID)
//			post.SetID(ID)
//			post.SetTitle(title)
//			post.SetContent(content)
//			post.SetImage(image)
//			post.SetUserId(userId)
//			post.SetCreatedAt(createdAt)
//			post.SetUpdatedAt(updatedAt)
//			post.SetOwner(owner.GetName())
//			post.SetCategories(cats)
//			post.SetComments(comments)
//			post.SetLike(likes)
//			post.SetDislike(dislikes)
//			post.SetComment(comment)
//			posts = append(posts, post)
//		}
//		return posts, nil
//	}
func (p *PostService) GetLastPostID() (uint, error) {
	var id uint
	err := p.GetDB().QueryRow(`SELECT MAX(id) FROM posts`).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (p *PostService) CountPosts() (uint, error) {
	rows, err := p.GetDB().Query(`SELECT * FROM posts`)
	if err != nil {
		return 0, err
	}
	var posts uint
	for rows.Next() {
		posts++
	}
	return posts, nil
}
func (p *PostService) GetPostsByCategory(categoryId uint) ([]models.Post, error) {
	rows, err := p.GetDB().Query(`SELECT post_id FROM category_post WHERE category_id = ?`, categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var (
			post   models.Post
			postID uint
		)
		err = rows.Scan(&postID)
		if err != nil {
			return nil, err
		}
		post, err = p.GetPost(postID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

//	func (p *PostService) GetLatestPosts() ([]models.Post, error) {
//		// Modifier la requête SQL pour trier par date de création en ordre décroissant et limiter à deux
//		rows, err := p.db.Query(`SELECT id, title, content, image, user_id, created_at FROM posts ORDER BY created_at DESC LIMIT 2`)
//		if err != nil {
//			return nil, err
//		}
//		defer rows.Close()
//		var posts []models.Post
//		for rows.Next() {
//			var (
//				post            models.Post
//				id              uint
//				title           string
//				content         string
//				image           string
//				userId          string
//				createdAt       time.Time
//				userService     UserService
//				categoryService CategoryService
//				commentService  CommentService
//				actionService   ActionService
//			)
//			userService.SetDB(p.GetDB())
//			categoryService.SetDB(p.GetDB())
//			commentService.SetDB(p.GetDB())
//			actionService.SetDB(p.GetDB())
//			err = rows.Scan(&id, &title, &content, &image, &userId, &createdAt)
//			if err != nil {
//				return nil, err
//			}
//			userID, err := pkg.StringToUUID(userId)
//			if err != nil {
//				return nil, err
//			}
//			owner, err := userService.GetUserById(userID)
//			if err != nil {
//				return nil, err
//			}
//			cats, err := categoryService.GetCategoriesByPost(id)
//			if err != nil {
//				return nil, err
//			}
//			comments, _ := commentService.GetCommentsByPosts(id)
//			comment, _ := commentService.CountCommentByPost(id)
//			likes, _ := actionService.GetLikesByPost(id)
//			dislikes, _ := actionService.GetDislikesByPost(id)
//			post.SetID(id)
//			post.SetTitle(title)
//			post.SetContent(content)
//			post.SetImage(image)
//			post.SetUserId(userID)
//			post.SetCreatedAt(createdAt)
//			post.SetOwner(owner.GetName())
//			post.SetCategories(cats)
//			post.SetComments(comments)
//			post.SetLike(likes)
//			post.SetDislike(dislikes)
//			post.SetComment(comment)
//			posts = append(posts, post)
//		}
//		return posts, nil
//	}
// func (p *PostService) GetLikedPostsByUser(userId uuid.UUID) ([]models.Post, error) {
// 	// var (
// 	// 	actionService ActionService
// 		post          models.Post
// 		posts         []models.Post
// 	)
// 	actionService.SetDB(p.GetDB())
// 	liked, err := actionService.GetLikedPostByUser(userId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, like := range liked {
// 		post, err = p.GetPost(like)
// 		if err != nil {
// 			return nil, err
// 		}
// 		posts = append(posts, post)
// 	}
// 	return posts, nil
// }
