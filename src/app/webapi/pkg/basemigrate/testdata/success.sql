--changeset josephspurrier:1
CREATE TABLE player (
    id BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    steam_id VARCHAR(191) NOT NULL,
    account_name VARCHAR(191) NOT NULL,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    PRIMARY KEY (id),
    CONSTRAINT u_player_account_name UNIQUE (account_name)
);
--rollback DROP TABLE IF EXISTS player;
--changeset josephspurrier:2
CREATE TABLE player_login (
    id BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    player_id BIGINT(10) UNSIGNED NOT NULL,
    display_name VARCHAR(191) NOT NULL,
    ip VARCHAR(191) NOT NULL,
    login_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY fk_player_login_player_id_player_id (player_id)
        REFERENCES player(id) ON DELETE CASCADE
);
--rollback DROP TABLE IF EXISTS player_login;

--changeset josephspurrier:3
CREATE TABLE game (
    id BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    start_at TIMESTAMP NULL,
    end_at TIMESTAMP NULL,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    PRIMARY KEY (id)
);
--rollback DROP TABLE IF EXISTS game;

--changeset josephspurrier:4
CREATE TABLE game_winner (
    id BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    game_id BIGINT(10) UNSIGNED NOT NULL,
    player_id BIGINT(10) UNSIGNED NOT NULL,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY fk_game_winner_game_id_game_id (game_id)
        REFERENCES game(id) ON DELETE CASCADE,
    FOREIGN KEY fk_game_winner_player_id_player_id (player_id)
        REFERENCES player(id) ON DELETE CASCADE
);
--rollback DROP TABLE IF EXISTS game_winner;

--changeset josephspurrier:5
CREATE TABLE game_player (
    id BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    game_id BIGINT(10) UNSIGNED NOT NULL,
    player_id BIGINT(10) UNSIGNED NOT NULL,
    kills SMALLINT(10) UNSIGNED NOT NULL,
    deaths SMALLINT(10) UNSIGNED NOT NULL,
    assists SMALLINT(10) UNSIGNED NOT NULL,
    damage_dealt MEDIUMINT(10) UNSIGNED NOT NULL,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY fk_game_player_game_id_game_id (game_id)
        REFERENCES game(id) ON DELETE CASCADE,
    FOREIGN KEY fk_game_player_player_id_player_id (player_id)
        REFERENCES player(id) ON DELETE CASCADE
);
--rollback DROP TABLE IF EXISTS game_player;