CREATE TABLE Users
(
    id       VARCHAR PRIMARY KEY NOT NULL,
    name     VARCHAR UNIQUE      NOT NULL,
    email    VARCHAR UNIQUE      NOT NULL,
    password VARCHAR             NOT NULL,
    role     INTEGER             NOT NULL
);

CREATE TABLE Contests
(
    id          VARCHAR PRIMARY KEY NOT NULL,
    title       VARCHAR             NOT NULL UNIQUE,
    description TEXT,
    startAt     timestamptz         NOT NULL,
    endAt       timestamptz         NOT NULL
);

CREATE TABLE Contestants
(
    id        VARCHAR PRIMARY KEY NOT NULL,
    role      INTEGER             NOT NULL,
    point     INTEGER             NOT NULL,

    contestID VARCHAR             NOT NULL,
    userID    VARCHAR             NOT NULL,

    FOREIGN KEY (contestID) REFERENCES Contests (id),
    FOREIGN KEY (userID) REFERENCES Users (id)
);

CREATE TABLE Problems
(
    id          VARCHAR PRIMARY KEY NOT NULL,
    index       VARCHAR             NOT NULL,
    title       VARCHAR             NOT NULL,
    text        TEXT                not null,
    point       INTEGER             NOT NULL,
    memoryLimit INTEGER             NOT NULL,
    timeLimit   INTEGER             NOT NULL,

    contestID   VARCHAR             NOT NULL,
    UNIQUE (index, title),
    FOREIGN KEY (contestID) REFERENCES Contests (id)
);

CREATE TABLE Casesets
(
    id        VARCHAR PRIMARY KEY NOT NULL,
    name      VARCHAR             NOT NULL,
    point     INTEGER             NOT NULL,

    problemID VARCHAR             NOT NULL,
    FOREIGN KEY (problemID) REFERENCES Problems (id)
);

CREATE TABLE Cases
(
    id        VARCHAR PRIMARY KEY NOT NULL,
    input     TEXT                NOT NULL,
    output    TEXT                NOT NULL,

    caseSetID VARCHAR             NOT NULL,
    FOREIGN KEY (caseSetID) REFERENCES Casesets (id)
);

CREATE TABLE Submissions
(
    id           VARCHAR PRIMARY KEY NOT NULL,
    point        INTEGER             NOT NULL,
    lang         VARCHAR             NOT NULL,
    codeLength   INTEGER             NOT NULL,
    result       VARCHAR             NOT NULL,
    execTime     INTEGER             NOT NULL,
    execMemory   INTEGER             NOT NULL,
    code         TEXT                NOT NULL,
    submittedAt  TIMESTAMPTZ         not null,

    problemID    VARCHAR             NOT NULL,
    contestantID VARCHAR             NOT NULL,
    FOREIGN KEY (problemID) REFERENCES Problems (id),
    FOREIGN KEY (contestantID) REFERENCES Contestants (id)
);

CREATE TABLE Submissionresults
(
    id           VARCHAR PRIMARY KEY NOT NULL,
    result       VARCHAR             NOT NULL,
    output       TEXT                NOT NULL,
    caseName     VARCHAR             NOT NULL,
    exitStatus   INTEGER             NOT NULL,
    execTime     INTEGER             NOT NULL,
    execMemory   INTEGER             NOT NULL,

    submissionID VARCHAR             NOT NULL,
    FOREIGN KEY (submissionID) REFERENCES Submissions (id)
)
