#!/bin/bash 
i="0"
while [ $i -lt 100 ]
do
	wget "http://www.bencinaenlinea.cl/web2/img/distribuidoras/logo$i.jpg" ../../assets/brand_icons
	i=$[$i+1]
done

