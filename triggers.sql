delimiter $$
DROP TRIGGER IF EXISTS after_comment_vote_insert;$$
CREATE TRIGGER after_comment_vote_insert
	AFTER INSERT ON comment_votes
	FOR EACH ROW
	BEGIN
		IF NEW.positive = TRUE THEN
			UPDATE comments SET votes = votes + 1 WHERE id = NEW.comment_id;
		ELSE
      UPDATE comments SET votes = votes - 1 WHERE id = NEW.comment_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_comment_vote_edit;$$
CREATE TRIGGER after_comment_vote_edit
	AFTER UPDATE ON comment_votes
	FOR EACH ROW
	BEGIN
		IF NEW.positive = TRUE THEN
			UPDATE comments SET votes = votes + 1 WHERE id = NEW.comment_id;
		ELSE
      UPDATE comments SET votes = votes - 1 WHERE id = NEW.comment_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_comment_vote_delete;$$
CREATE TRIGGER after_comment_vote_delete
	AFTER DELETE ON comment_votes
	FOR EACH ROW
	BEGIN
		IF OLD.positive = TRUE THEN
			UPDATE comments SET votes = votes - 1 WHERE id = OLD.comment_id;
		ELSE
      UPDATE comments SET votes = votes + 1 WHERE id = OLD.comment_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_post_vote_insert;$$
CREATE TRIGGER after_post_vote_insert
	AFTER INSERT ON post_votes
	FOR EACH ROW
	BEGIN
		IF NEW.positive = TRUE THEN
			UPDATE posts SET votes = votes + 1 WHERE id = NEW.post_id;
		ELSE
      UPDATE posts SET votes = votes - 1 WHERE id = NEW.post_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_post_vote_edit;$$
CREATE TRIGGER after_post_vote_edit
	AFTER UPDATE ON post_votes
	FOR EACH ROW
	BEGIN
		IF NEW.positive = TRUE THEN
			UPDATE posts SET votes = votes + 1 WHERE id = NEW.post_id;
		ELSE
      UPDATE posts SET votes = votes - 1 WHERE id = NEW.post_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_post_vote_delete;$$
CREATE TRIGGER after_post_vote_delete
	AFTER DELETE ON post_votes
	FOR EACH ROW
	BEGIN
		IF OLD.positive = TRUE THEN
			UPDATE posts SET votes = votes - 1 WHERE id = OLD.post_id;
		ELSE
      UPDATE posts SET votes = votes + 1 WHERE id = OLD.post_id;
		END IF;
	END ;$$


