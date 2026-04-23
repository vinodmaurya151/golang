1. Initialize Git (if not already done)
    git init // Translation: "Hey Git, start watching this folder"


2. Add your files to staging
    # Add specific file
    git add filename.txt

    # Add all files(Translation: "Hey Git, look at ALL files in this folder")
    git add .

3. Commit your changes(Translation: "Take a photo/savepoint of all these files. Label this photo 'first commit'")
    git commit -m "Your commit message"  

4. Connect to remote repository
    # Add remote origin (replace with your repo URL)(Translation: "Hey Git, there's a cloud storage location at this address. Give it the nickname 'origin'")
    git remote add origin https://github.com/username/repository.git 

    # If you already have a remote, check it
    git remote -v

5. Push to GitHub/GitLab/Bitbucket(Translation: "Send my 'main' storyline to the cloud location called 'origin'")
    # First push (set upstream)
    git push -u origin main

    # Or if using master branch
    git push -u origin master

    # Subsequent pushes
    git push




Complete Example:
# Step-by-step workflow
cd your-project-folder
git init
git add .
git commit -m "Initial commit"
git remote add origin https://github.com/yourusername/your-repo.git
git branch -M main  # Rename branch to main
git push -u origin main



Quickest solution:

Just run this to remove the nested git and proceed:
bash

rmdir /s /q VM_21_04_2026\.git
git add .
git commit -m "Initial commit"
git remote add origin <your-repo-url>
git push -u origin main



#### Start Fresh ####
# Initialize Git
git init

# Add and commit
git add main.go
git commit -m "first commit"

# Check branch name
git branch
# It will show "* master"

# Create new repo on GitHub called "go-demo"
# Then push
git remote add origin https://github.com/vinodmaurya151/go-demo.git
git push -u origin master



###
Step 1: Create a new repository from git website
Step 2: Go to main folder(Go->1_print), Go is the main folder
Step 3: git init
Step 4: git add .
Step 5: git commit -m "1_print commit"
Step 6: git branch -M main
Step 7: git remote add origin https://github.com/vinodmaurya151/golang.git
Step 8: git push -u origin main
