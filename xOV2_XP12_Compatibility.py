import os
import subprocess

xp12_path = "../X-Plane.exe"
xp12_gs_folder = "SCENERY_PACK *GLOBAL_AIRPORTS*"
xo_dummy_folder = "SCENERY_PACK Custom Scenery/###GLOBAL SCENERY####/"

file1 = open('scenery_packs.ini', 'r')
Lines = file1.readlines()
file1.close()
#print(Lines)

index = Lines.index(xo_dummy_folder + "\n")
Lines[index] = xp12_gs_folder + "\n"


os.remove('scenery_packs.ini')

file1 = open('scenery_packs.ini', 'w')
for element in Lines:
    file1.write(element)
file1.close()

subprocess.Popen(xp12_path)