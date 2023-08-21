DROP DATABASE IF EXISTS simple_tiktok;
CREATE DATABASE simple_tiktok;
USE simple_tiktok;
CREATE TABLE `follow` (
  `id` integer PRIMARY KEY,
  `following_user_id` integer,
  `followed_user_id` integer
);

CREATE TABLE `user` (
  `id` integer PRIMARY KEY,
  `username` varchar(32),
  `password` varchar(34),
  `follow_count` bigint,
  `follower_count` bigint,
  `avatar` varchar(255),
  `background_image` varchar(255),
  `signature` text,
  total_favorite bigint,
  `work_count` bigint,
  `favorite_count` bigint
);

CREATE TABLE `video` (
  `id` integer PRIMARY KEY,
  `user_id` integer,
  `play_url` varchar(255),
  `cover_url` varchar(255),
  `title` varchar(255),
  `created_at` timestamp,
  `favorite_count` bigint,
  `comment_count` bigint
);

CREATE TABLE `comment` (
  `id` integer PRIMARY KEY,
  `user_id` integer,
  `content` text,
  `created_at` timestamp,
  `vedio_id` integer
);

CREATE TABLE `video_like` (
  `id` integer PRIMARY KEY,
  `user_id` integer,
  `video_id` integer
);

CREATE TABLE `chat` (
  `id` integer PRIMARY KEY,
  `user_id` integer,
  `content` text,
  `created_at` timestamp,
  `object` integer
);

CREATE TABLE `friend` (
  `id` integer PRIMARY KEY,
  `user_id` integer,
  `friend_id` integer
);

ALTER TABLE `follow` ADD FOREIGN KEY (`following_user_id`) REFERENCES `user` (`id`);

ALTER TABLE `follow` ADD FOREIGN KEY (`followed_user_id`) REFERENCES `user` (`id`);

ALTER TABLE `video` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `comment` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `comment` ADD FOREIGN KEY (`vedio_id`) REFERENCES `video` (`id`);

ALTER TABLE `video_like` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `video_like` ADD FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);

ALTER TABLE `chat` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `chat` ADD FOREIGN KEY (`object`) REFERENCES `user` (`id`);

ALTER TABLE `friend` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `friend` ADD FOREIGN KEY (`friend_id`) REFERENCES `user` (`id`);
