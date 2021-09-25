CREATE TABLE Users (
	UserId SERIAL PRIMARY KEY,
	Username TEXT
);

CREATE TABLE Posts (
    PostId SERIAL PRIMARY KEY,
    UserId INT NOT NULL,
    Title TEXT,
    CONSTRAINT fk_User FOREIGN KEY(UserId) REFERENCES Users(UserId) ON DELETE CASCADE
);

CREATE TABLE UserLikes (
    FollowedId INT,
    FollowerId INT,
    PRIMARY KEY (FollowedId, FollowerId),
    CONSTRAINT fk_Followed FOREIGN KEY(FollowedId) REFERENCES Users(UserId) ON DELETE CASCADE,
    CONSTRAINT fk_Follower FOREIGN KEY(FollowerId) REFERENCES Users(UserId) ON DELETE CASCADE
);

CREATE TABLE PostsLikes (
    UserId INT,
    PostId INT,
    PRIMARY KEY (UserId, PostId),
    CONSTRAINT fk_User FOREIGN KEY(UserId) REFERENCES Users(UserId) ON DELETE CASCADE,
    CONSTRAINT fk_Post FOREIGN KEY(PostId) REFERENCES Posts(PostId) ON DELETE CASCADE
);

INSERT INTO 
    Users (Username) 
VALUES 
    ('John Doe'),
    ('Mike Mock'),
    ('Lucy Applegate'),
    ('Sam Taylor');

INSERT INTO
    Posts (UserId, Title)
VALUES
    (1, 'lorem'),
    (1, 'ipsum'),
    (2, 'dolor'),
    (3, 'sit'),
    (3, 'amet'),
    (3, 'consectetur'),
    (4, 'adipiscing');

INSERT INTO
    UserLikes (FollowedId, FollowerId)
VALUES
    (3, 4),
    (3, 2),
    (1, 3),
    (3, 1),
    (1, 2),
    (2, 4);

INSERT INTO
    PostsLikes (UserId, PostId)
VALUES
    (1, 3),
    (2, 4),
    (3, 7),
    (4, 5);