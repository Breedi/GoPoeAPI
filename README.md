# GoPoeAPI

This is a API fetcher written in GO to get the current Trade API Data of Path of Exile. It summarizes everything inside the prepared data structure [TradeData]. 

## How does it work?

First of all. The API wont give you all available stashes at once. Just around 600. If you fetch your data the first time or with the default value 0 you will get a random NextChangeID. Inside your TradeData will be the ChangeID for the next "page" of data and your given stashes. You will be able to go through it and use it for your projects.

## Data Structure

| Data Structure Name  | Description  |
|---|---|
| TradeData  | NextChangeID plus all stashes from some player  |
| Stashes  |  Stash of some Player with Items and metadata |
| Items  | Information about given Items like Itemlevel, if it is corrupted etc |
| Properties  | Properties of an item like attackspeed values etc |
| Sockets  | Data strucutre for sockets / gems |
| Category  | Which type of item (gloves, helmet or similar) |



### Official Documentation: https://www.pathofexile.com/developer/docs/api-resource-public-stash-tabs


Feel Free to use it within your own projects, I try to develop it further.
