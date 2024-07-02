package database

var (
    // Création des tables
    createUsersTable = `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        nickname TEXT NOT NULL UNIQUE,
        age INTEGER,
        gender TEXT,
        firstname TEXT,
        lastname TEXT,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        image TEXT,
        active BOOLEAN,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

    createPostsTable = `CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        user TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        image TEXT,  
        room_id INTEGER
    );`

    createCommentsTable = `CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL,
        user TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        post_id INTEGER,
        FOREIGN KEY(post_id) REFERENCES posts(id)
    );`

    createActionsTable = `CREATE TABLE IF NOT EXISTS actions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user TEXT,
        post_id INTEGER,
        comment_id INTEGER,
        action TEXT NOT NULL
    );`

    createCategoriesTable = `CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`

    createMessagesTable = `CREATE TABLE IF NOT EXISTS messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL,
        user_send INTEGER,
        time DATETIME DEFAULT CURRENT_TIMESTAMP,
        room TEXT
    );`

    // Insertion de données
    insertUser = `INSERT INTO users (nickname, age, gender, firstname, lastname, email, password, image, active) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`
    insertPost = `INSERT INTO posts (title, content, user, created_at, updated_at, image, room_id) VALUES (?, ?, ?, ?, ?, ?, ?);`
    insertComment = `INSERT INTO comments (content, user, created_at, updated_at, post_id) VALUES (?, ?, ?, ?, ?);`
    insertAction = `INSERT INTO actions (user, post_id, comment_id, action) VALUES (?, ?, ?, ?);`
    insertCategory = `INSERT INTO categories (name) VALUES (?);`
    insertMessage = `INSERT INTO messages (content, user_send, time, room) VALUES (?, ?, ?, ?);`

    // Sélection de données
    getUserById = `SELECT id, nickname, age, gender, firstname, lastname, email, password, image, active, created_at FROM users WHERE id = ?;`
    getUserByEmail = `SELECT id, nickname, age, gender, firstname, lastname, email, password, image, active, created_at FROM users WHERE email = ?;`

    getPostById = `SELECT id, title, content, user, created_at, updated_at, image, room_id FROM posts WHERE id = ?;`
    getPosts = `SELECT id, title, content, user, created_at, updated_at, image, room_id FROM posts ORDER BY created_at DESC;`
    getPostsByUser = `SELECT id, title, content, user, created_at, updated_at, image, room_id FROM posts WHERE user = ? ORDER BY created_at DESC;`

    getCommentById = `SELECT id, content, user, created_at, updated_at, post_id FROM comments WHERE id = ?;`
    getCommentsByPostId = `SELECT id, content, user, created_at, updated_at, post_id FROM comments WHERE post_id = ? ORDER BY created_at DESC;`
    getCommentsByUser = `SELECT id, content, user, created_at, updated_at, post_id FROM comments WHERE user = ? ORDER BY created_at DESC;`

    getActionById = `SELECT id, user, post_id, comment_id, action FROM actions WHERE id = ?;`
    getActionsByUser = `SELECT id, user, post_id, comment_id, action FROM actions WHERE user = ? ORDER BY id DESC;`

    getCategoryById = `SELECT id, name FROM categories WHERE id = ?;`
    getCategories = `SELECT id, name FROM categories ORDER BY id;`

    getMessageById = `SELECT id, content, user_send, time, room FROM messages WHERE id = ?;`
    getMessagesByRoom = `SELECT id, content, user_send, time, room FROM messages WHERE room = ? ORDER BY time DESC;`

    // Mise à jour de données
    updateUser = `UPDATE users SET nickname = ?, age = ?, gender = ?, firstname = ?, lastname = ?, email = ?, password = ?, image = ?, active = ? WHERE id = ?;`
    updatePost = `UPDATE posts SET title = ?, content = ?, user = ?, updated_at = ?, image = ?, room_id = ? WHERE id = ?;`
    updateComment = `UPDATE comments SET content = ?, user = ?, updated_at = ? WHERE id = ?;`

    // Suppression de données
    deleteUser = `DELETE FROM users WHERE id = ?;`
    deletePost = `DELETE FROM posts WHERE id = ?;`
    deleteComment = `DELETE FROM comments WHERE id = ?;`
    deleteAction = `DELETE FROM actions WHERE id = ?;`
    deleteCategory = `DELETE FROM categories WHERE id = ?;`
    deleteMessage = `DELETE FROM messages WHERE id = ?;`

    // Authentification
    authenticateUser = `SELECT id, nickname, password FROM users WHERE nickname = ?;`

    // Recherche
    searchPostsByTitle = `SELECT id, title, content, user, created_at, updated_at, image, room_id FROM posts WHERE title LIKE ? ORDER BY created_at DESC;`
    searchPostsByContent = `SELECT id, title, content, user, created_at, updated_at, image, room_id FROM posts WHERE content LIKE ? ORDER BY created_at DESC;`
)
