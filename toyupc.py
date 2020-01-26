import pandas as pd
import requests
from bs4 import BeautifulSoup
import numpy as np

BASE_URL = [
'653569825586',
#'http://www.reuters.com/finance/stocks/company-officers/AMZN',
#'http://www.reuters.com/finance/stocks/company-officers/AAPL'
]


#loading empty array for toys
toys = []
#Loop through our URLs we loaded above
for b in BASE_URL:
    toyname=''
    manufacturer=''
    html = requests.get('https://www.upcitemdb.com/upc/' + b).text
    soup = BeautifulSoup(html, "html.parser")
    #identify table we want to scrape
    name_list = soup.find('ol',{"class":"num"})
    ict=0
    for namerow in name_list.find_all('li'):
        ict +=1
        if ict==1:
            toyname=namerow.text.strip()

    detail_table = soup.find('table', {"class" : "detail-list"})

    #try clause to skip any companies with missing/empty board member tables
    try:
    #loop through table, grab each of the 4 columns shown (try one of the links yourself to see the layout)
        for row in detail_table.find_all('tr'):
            cols = row.find_all('td')
            if cols[0].text.strip()=="Brand:":
                manufacturer=cols[1].text.strip()
    except: pass  
    toys.append((b,toyname,manufacturer))

toy_array = np.asarray(toys)
df = pd.DataFrame(toy_array)
df.columns = ['UPC', 'Toy Name', 'Manufacturer']
print(df)

#board_array = np.asarray(board_members)
#print(len(board_array))
#df = pd.DataFrame(board_array)
#df.columns = ['URL', 'Name', 'Age','Year_Joined', 'Title']
#print(df.head(10))