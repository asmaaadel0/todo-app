CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "title" TEXT,
    "completed" boolean
);

SELECT * FROM tasks;

INSERT INTO tasks(title, completed) VALUES (?, ?);

DELETE FROM tasks WHERE id = ?;

UPDATE tasks SET title = ?, completed = ? WHERE id = ?;