## README.md

I wrote this to estimate the cost of prepaid tokens for domestic consumers on the KPLC DC tariff prior to paying for them. This was in response to an inquiry on  [skunkworks](http://orion.my.co.ke/cgi-bin/mailman/listinfo/skunkworks), a mailing list frequented by kenyan techies.  

Runs on [Google App Engine.](https://developers.google.com/appengine/) 
 
The various charges and levies are obtained from public information available in the 
[Kenya gazette](http://kenyalaw.org/kenya_gazette/). Please note that KPLC updates the official rates in the Kenya gazette towards end month (if at all) and therefore the rates used will lag behind by a month or so.


##TODO

1. Clean and comment the code.
2. Handle second and subsequent purchases.
3. Handle purchases where the last purchase is more than a month removed from the current month. This involves obtaining month of last purchase and fixed charges heretofore accrued.



##You might want to know
I did this as a learning project for the [Go](http://golang.org) programming language. Thus the code itself is at a beginner level thus it is quite easy to follow in most parts, and needlessly convoluted in some.  Fork and upgrade it!

##License

Licensed under [GPL v3](http://www.gnu.org/licenses/gpl-3.0.txt) (the "License"), you may not use this program except in compliance with the License. You may obtain a copy of the License at

	http://www.gnu.org/licenses/gpl-3.0.txt
	
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License version 3, or (at your option) any later version for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

TL;DR [click here](http://choosealicense.com/licenses/#gpl-v3)
 