#!/bin/bash

# Menambahkan semua perubahan ke staging area
git add .

# Membuat commit dengan pesan "first commit"
git commit -m "first commit"

# Mengirim perubahan ke branch master di remote origin
git push origin master