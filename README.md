# splendor
璀璨宝石服务端/客户端，练手项目

## 开发进度

1. 完成牌堆打乱【DONE】
   1. 发展牌打乱
   2. 贵族卡打乱
2. 完成分发发展牌，分发【DONE】
   1. 发展牌每个级别分4张
   2. 贵族卡抽人数+1分发
3. 完成玩家购牌系统
   1. 完成购牌，并且自动补充卡牌【DONE】
   2. 完成购牌后消费宝石，并将发展牌加入，并增加奖励
   3. 购买自己的手牌
   4. 购买发展卡是统计上奖励
   5. 选择保留一张卡牌，并且获取一个黄金
   6. 实现使用黄金购买
   7. 摸取一张牌堆，并保存，设置保留上限


## Content

Components
- 7 Emerald tokens (green)
- 7 Diamond tokens (white)
- 7 Sapphire tokens (blue)
- 7 Onyx tokens(black)
- 7 Ruby tokens (red)
- 5 Gold Joker tokens (yellow)
- 90 Development cards
- 10 Noble tiles

前面6个就是宝石，然后是发展卡，然后是贵族卡

## Setup

Shuffle each development card deck separately, and then place them in a column in the middle of the table in increasing order from bottom to top (level 1;level 2;level 3).
> 分别洗牌每个发展卡组，然后将它们按从下到上的递增顺序放在表格中间的一列中（级别1；级别2；级别3）。

Then reveal 4 cards from each level.
> 然后展示每个级别的4张卡片。

Shuffle the noble tiles and reveal as many of them as there are players plus one (example: 5 tiles for a 4 player game).
> 打乱贵族牌，并显示出玩家加一的数量（例如：4人游戏5个牌）。

The remaining tiles are removed from the game; they will not be used during the game. Finally, place the tokens in 6 distinct piles (sort them by color) within reach of the players.
> 剩余的贵族牌将从游戏中移除；游戏期间不会使用它们。最后，将宝石放置在玩家触手可及的6个不同的堆中（按颜色排序）。

- Game With 2 Players
  - Remove 3 tokens of each gem color (there should be only 4 of each remaining).
  - Don't touch the gold.
  - Reveal 3 noble tiles.

- Game With 3 Players
  - Remove 2 tokens of each gem color (there should be only 5 of each remaining).
  - Don't touch the gold.
  - Reveal 4 noble tiles.

## The Development Cards
To win prestige points, the players must purchase development cards. These cards are visible in the middle of the table and may be purchased by all players during the game.
> 为了赢得声望积分，玩家必须购买发展卡。这些卡片在桌子中间可见，所有玩家都可以在游戏期间购买。

The developments in hand are the cards which the players reserve throughout the game. Developments in hand may only be purchased by the players holding them.
> 手中的牌是玩家在整个游戏中保留的牌。手中的发展卡只能由持有发展卡的玩家购买。


![发展卡展示](/pics/img_1.png)

> Example: The player who purchases this card wins 4 prestige points.
> 
> 购买此卡的玩家将获得4点声望。
> 
> Owning this card allows the player to benefit from a blue bonus. To purchase this card, the player must spend 3 blue tokens, 3 black tokens, and 6 white tokens.
> 
> 拥有这张卡可以让玩家从蓝色奖金中受益。要购买此卡，玩家必须花费3个蓝色代币、3个黑色代币和6个白色代币。

## The Noble Tiles
The noble tiles are visible in the middle of the table. At the end of their turn, a player automatically receives the visit from a noble if that player has the amount of bonuses (and only bonuses) required, and they get the corresponding tile.
> 贵族卡在桌子中间可见。在他们的回合结束时，如果玩家有所需的奖金（并且只有奖金），玩家将自动接受贵族的访问，并获得相应的牌。

A player cannot refuse a visit from a noble.
> 玩家不能拒绝贵族的来访。

Receiving a noble isn't considered to be an action. Each noble tile is worth 3 prestige points, but players can only get a single one per turn.
> 接待贵族不被认为是一种行为。每个贵族牌值3点声望值，但玩家每回合只能获得一个。

![发展卡展示](/pics/img_3.png)

> Example: The player who acquires this noble's tile gains 3 prestige points.
>
> 获得此贵族牌的玩家获得3点声望。
> 
> If a player has enough development cards to amass 3 blue bonuses, 3 green bonuses, and 3 white bonuses, then that player automatically receives a visit from that noble.
> 
> 如果一个玩家有足够的发展卡来累积3个蓝色奖金、3个绿色奖金和3个白色奖金，那么该玩家将自动获得该贵族的访问。

## Game Play
The youngest player begins. Play then proceeds clockwise.
> 最年轻的玩家开始比赛。然后顺时针播放。

On their turn, a player must choose to perform only one of the following four actions:
> 轮到他们时，玩家必须选择执行以下四个动作中的一个：

- Take 3 gem tokens of different colors.
- > 取3个不同颜色的宝石代币。

- Take 2 gem tokens of the same color. This action is only possible if there are at least 4 tokens of the chosen color left when the player takes them.
- > 取2个相同颜色的宝石代币。只有当玩家拿走所选颜色的代币时，至少剩下4个代币时，才能执行此操作。
  > 不能剩下一个，意思就是3个的时候不能执行这个操作。

- Reserve 1 development card and take 1 gold token (joker).
- > 保留1张发展卡并获得1个金币（joker）。

- Purchase 1 face-up development card from the middle of the table or a previously reserved one.
- > 从桌子中间购买1张面朝上的发展卡或之前预留的发展卡。

## Selecting Tokens
A player can never have more than 10 tokens at the end of their turn (including jokers). If this happens, they must return tokens until they only have 10 left. A player can return all or some of those they've just drawn. The tokens owned by a player must be visible by all players at all times.
> 一名玩家在回合结束时的代币数不能超过10个（包括黄金）。如果发生这种情况，他们必须返回令牌，直到只剩下10个为止。玩家可以退还他们刚刚抽到的全部或部分。玩家拥有的代币必须在任何时候都能被所有玩家看到。

> Note: players may not take 2 tokens of the same color if there are less than 4 tokens available of that color.
> 
> 如果有少于4个相同颜色的代币可用，则玩家可能无法获得2个相同颜色代币。

## Reserve a development card
To reserve a card, a player simply needs to take a face-up development from the middle of the table or (if you're feeling lucky) draw the first card from one of the three decks (level 1; level 2 ; level 3) without showing it to the other players.
> 要保留一张牌，玩家只需从桌子中间进行正面朝上的发展，或者（如果你觉得幸运）从三张牌中的一张牌（1级；2级；3级）中抽出第一张牌，而不向其他玩家展示。

The reserved cards are kept in hand and cannot be discarded. Players may not have more than three reserved cards in hand, and the only way to get rid of a card is to buy it.
> 保留的卡保存在手中，不能丢弃。玩家手中的保留卡可能不超过三张，而摆脱一张卡的唯一方法就是购买它。

Reserving a card is also the only way to get a gold token (joker). If there is no gold left, you can still reserve a card, but you won't get any gold.
> 预订一张卡也是获得金币的唯一方式（joker）。如果没有剩余的黄金，你仍然可以保留一张卡，但你不会得到任何黄金。

## Buying a development card
To purchase a card, a player must spend the number of tokens indicated on the card. A joker token can replace any color. The spent tokens (including any jokers) are returned to the middle of the table.
> 要购买一张卡，玩家必须花费卡上显示的宝石数量。黄金可以替换任何颜色。花费的宝石（包括任何黄金）将返回到桌子中间。

A player may purchase one of the face-up development cards in the middle of the table or a card in his hand that was reserved on a previous turn.
> 玩家可以购买位于桌子中间的一张面朝上的发展卡，也可以购买他手中在前一回合中保留的卡。

Each player makes distinct rows with the acquired development cards by sorting them by color, and staggering them vertically so that their bonuses and prestige point values are visible.
> 每个玩家通过按颜色对获得的发展卡进行排序，并将其垂直错开，以使其奖金和声望值可见，从而使其与获得的发展卡片形成不同的行。

The bonuses and prestige points granted by each card must be visible to all at all times.
> 每张卡所授予的奖金和声望点数必须始终可见。

Note: when a development card from the middle of the table is acquired or reserved, it must immediately be replaced by a card of the same level.
> 注意：当获取或保留表中间的发展卡时，必须立即用相同级别的卡替换。

At all times during the game, there must be 4 face-up cards of each level (unless the deck in question is empty, in which case the empty spaces also remain empty).
> 在游戏期间的任何时候，每个级别都必须有4张面朝上的牌（除非所讨论的牌组是空的，在这种情况下，空白区域也保持空白）。

## The Bonuses
The bonuses a player has from development cards acquired on previous turns provide discounts on the purchase of new cards. Each bonus of a given color is equal to a token of that color.
> 玩家从先前回合获得的发展卡获得的奖金为购买新卡提供折扣。给定颜色的每个加成等于该颜色的标记。

Thus, if a player has 2 blue bonuses and wants to purchase a card which costs 2 blue tokens and 1 green token, the player must only spend 1 green token.
> 因此，如果玩家有2个蓝色奖金，并且想要购买一张花费2个蓝色代币和1个绿色代币的卡，则玩家只能花费1个绿色令牌。

If a player has enough development cards (and therefore bonuses), they can even purchase a card without spending any tokens.
> 如果玩家有足够的发展卡（因此奖金），他们甚至可以购买一张卡，而不需要花费任何代币。

## The Nobles
At the end of their turn, each player checks the noble tiles in order to determine if they're receiving a visit from one of them. A player can be visited if they have (at least) the quantity and type of bonuses indicated on the noble tile.
> 在他们的回合结束时，每个玩家检查贵族牌，以确定他们是否接受其中一个玩家的访问。如果玩家拥有（至少）贵族牌上所示的奖金数量和类型，则可以访问该玩家。

It is impossible to refuse the visit from a noble, which is not considered to be an action.
> 拒绝贵族的来访是不可能的，这不被认为是一种行为。

If a player has enough bonuses to be visited by more than one noble at the end of their turn, that player chooses the noble to be received.
> 如果一名玩家在回合结束时有足够的奖金可供多名贵族访问，则该玩家选择要获得的贵族。

The tile obtained is placed face-up in front of the player in question.
> 获得的平铺面朝上放置在所讨论的玩家面前。


## End of the Game

When a player reaches 15 prestige points, complete the current round so that each player has played the same number of turns.
> 当一名玩家达到15个声望点时，完成当前回合，使每个玩家玩了相同的回合数。

The player who then has the highest number of prestige points is declared the winner (don't forget to count your nobles). In case of a tie, the player who has purchased the fewest development cards wins.
> 获得最高声望点数的玩家被宣布为获胜者（别忘了计算你的贵族）。在平局的情况下，购买最少发展卡的玩家获胜。


## Reference
- https://www.ultraboardgames.com/splendor/game-rules.php
- https://www.wikihow.com/Play-Splendor