delimiter $$
DROP TRIGGER IF EXISTS after_comment_vote_insert;$$
CREATE TRIGGER after_comment_vote_insert
	AFTER INSERT ON comment_votes
	FOR EACH ROW
	BEGIN
		IF NEW.positive = TRUE THEN
			UPDATE comments SET votes = votes + 1 WHERE comment_id = NEW.comment_id;
		ELSE
      UPDATE comments SET votes = votes - 1 WHERE comment_id = NEW.comment_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_comment_vote_edit;$$
CREATE TRIGGER after_comment_vote_edit
	AFTER UPDATE ON comment_votes
	FOR EACH ROW
	BEGIN
		IF NEW.positive = TRUE THEN
			UPDATE comments SET votes = votes + 1 WHERE comment_id = NEW.comment_id;
		ELSE
      UPDATE comments SET votes = votes - 1 WHERE comment_id = NEW.comment_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_comment_vote_delete;$$
CREATE TRIGGER after_comment_vote_delete
	AFTER DELETE ON comment_votes
	FOR EACH ROW
	BEGIN
		IF OLD.positive = TRUE THEN
			UPDATE comments SET votes = votes - 1 WHERE comment_id = OLD.comment_id;
		ELSE
      UPDATE comments SET votes = votes + 1 WHERE comment_id = OLD.comment_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_post_vote_insert;$$
CREATE TRIGGER after_post_vote_insert
	AFTER INSERT ON post_votes
	FOR EACH ROW
	BEGIN
		IF NEW.positive = TRUE THEN
			UPDATE posts SET votes = votes + 1 WHERE post_id = NEW.post_id;
		ELSE
      UPDATE posts SET votes = votes - 1 WHERE post_id = NEW.post_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_post_vote_edit;$$
CREATE TRIGGER after_post_vote_edit
	AFTER UPDATE ON post_votes
	FOR EACH ROW
	BEGIN
		IF NEW.positive = TRUE THEN
			UPDATE posts SET votes = votes + 1 WHERE post_id = NEW.post_id;
		ELSE
      UPDATE posts SET votes = votes - 1 WHERE post_id = NEW.post_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_post_vote_delete;$$
CREATE TRIGGER after_post_vote_delete
	AFTER DELETE ON post_votes
	FOR EACH ROW
	BEGIN
		IF OLD.positive = TRUE THEN
			UPDATE posts SET votes = votes - 1 WHERE post_id = OLD.post_id;
		ELSE
      UPDATE posts SET votes = votes + 1 WHERE post_id = OLD.post_id;
		END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_comment_add;$$
CREATE TRIGGER after_comment_add
	AFTER INSERT ON comments
	FOR EACH ROW
	BEGIN
		UPDATE posts set comment_quantity = comment_quantity + 1 WHERE id = NEW.post_id;
        IF NEW.father != 0 THEN
			UPDATE comments set comment_quantity = comment_quantity + 1 WHERE id = NEW.father;
        END IF;
	END ;$$

DROP TRIGGER IF EXISTS after_comment_delete;$$
CREATE TRIGGER after_comment_delete
	AFTER DELETE ON comments
	FOR EACH ROW
	BEGIN
		UPDATE posts set comment_quantity = comment_quantity - 1 WHERE id = OLD.post_id;
        IF OLD.father != 0 THEN
			UPDATE comments set comment_quantity = comment_quantity - 1 WHERE id = OLD.father;
        END IF;
	END ;$$