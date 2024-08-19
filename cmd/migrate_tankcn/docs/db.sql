alter table wt_player
    AUTO_INCREMENT = 30000000;

alter table wt_legion
    AUTO_INCREMENT = 30000;

alter table wt_gift
    AUTO_INCREMENT = 400000;

DELETE
FROM wt_gamecenter
WHERE player_id >= 20000000
  AND player_id < 30000000;

DELETE
FROM wt_legion
WHERE id >= 20000
  AND id < 30000;

DELETE
FROM wt_player
WHERE player_id >= 20000000
  AND player_id < 30000000;
