/**
 * Initialize your data structure here.
 */
var Twitter = function() {
    // <follower, followee[]>
    // posts
    this.followees = new Map();
    this.posts = [];
};

var Post = function(tweetId, userId) {
    this.tweetId = tweetId;
    this.userId = userId;
};

/**
 * Compose a new tweet. 
 * @param {number} userId 
 * @param {number} tweetId
 * @return {void}
 */
Twitter.prototype.postTweet = function(userId, tweetId) {
    this.posts.push(new Post(tweetId, userId));
};

/**
 * Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. 
 * @param {number} userId
 * @return {number[]}
 */
Twitter.prototype.getNewsFeed = function(userId) {
    let followees = this.followees.get(userId);
    if (followees === undefined) {
        followees = [];
    }

    let posts = [];
    for (let i = this.posts.length-1; i >= 0 ; i--) {
        let post = this.posts[i];
        if (post.userId === userId) {
            posts.push(post.tweetId);
        }
        for (let j = 0; j < followees.length; j++) {
            if (post.userId === followees[j]) {
                posts.push(post.tweetId);
                break;
            }
        }
        if (posts.length === 10) {
            return posts;
        }
    }

    return posts;
};

/**
 * Follower follows a followee. If the operation is invalid, it should be a no-op. 
 * @param {number} followerId 
 * @param {number} followeeId
 * @return {void}
 */
Twitter.prototype.follow = function(followerId, followeeId) {
    // check if invalid
    if (followerId === followeeId) {
        return;
    }

    // get current followees of this follower
    let followees = this.followees.get(followerId);
    if (followees === undefined) {
        followees = [];
    }

    // check if valid op
    for (var i = 0; i < followees.length; i++) {
        if (followees[i] === followeeId) {
            break;
        }
    }
    // insert new followee
    if (i === followees.length) {
        followees.push(followeeId);
        this.followees.set(followerId, followees);
        return;
    }
};

/**
 * Follower unfollows a followee. If the operation is invalid, it should be a no-op. 
 * @param {number} followerId 
 * @param {number} followeeId
 * @return {void}
 */
Twitter.prototype.unfollow = function(followerId, followeeId) {
    // check if invalid
    if (followerId === followeeId) {
        return;
    }

    // get current followees of this follower
    let followees = this.followees.get(followerId);
    if (followees === undefined) {
        followees = [];
    }

    // check if valid op
    for (var i = 0; i < followees.length; i++) {
        if (followees[i] === followeeId) {
            break;
        }
    }
    // remove new followee
    if (i !== followees.length) {
        followees = followees.slice(0, i).concat(followees.slice(i+1));
        this.followees.set(followerId, followees);
        return;
    }
};

/** 
 * Your Twitter object will be instantiated and called as such:
 * var obj = new Twitter()
 * obj.postTweet(userId,tweetId)
 * var param_2 = obj.getNewsFeed(userId)
 * obj.follow(followerId,followeeId)
 * obj.unfollow(followerId,followeeId)
 */

let twitter = new Twitter();
twitter.postTweet(1, 5);
console.log(twitter.getNewsFeed(1));
twitter.follow(1, 2);
twitter.postTweet(2, 6);
console.log(twitter.getNewsFeed(1));
twitter.unfollow(1, 2);
console.log(twitter.getNewsFeed(1));