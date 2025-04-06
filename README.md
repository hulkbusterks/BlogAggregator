# Blog Aggregator CLI

This command-line tool, `BlogAggregator`, allows users to manage and aggregate content from RSS feeds. It provides functionalities to register users, add and follow feeds, and browse collected posts.

## Prerequisites

-   Go (Golang) installed.
-   PostgreSQL database set up and running.
-   `lib/pq` PostgreSQL driver installed.

## Installation and Setup

1.  **Clone the Repository:**

    ```bash
    git clone <repository_url>
    cd BlogAggregator
    ```

2.  **Install Dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Configure Database:**

    -   Create a `.env` file in the root directory with the following content:

        ```
        DB_URL="postgres://<user>:<password>@<host>:<port>/<dbname>?sslmode=disable"
        ```

        Replace `<user>`, `<password>`, `<host>`, `<port>`, and `<dbname>` with your PostgreSQL database credentials.

4.  **Build the Application:**

    ```bash
    go build -o BlogAggregator
    ```

## Usage

### Reset Database

-   Resets the database, clearing all data.

    ```bash
    ./BlogAggregator reset
    ```

### Register User

-   Registers a new user.

    ```bash
    ./BlogAggregator register <username>
    ```

    Example:

    ```bash
    ./BlogAggregator register abc
    ```

### Add Feed

-   Adds a new RSS feed to the system.

    ```bash
    ./BlogAggregator addfeed "<feed_name>" "<feed_url>"
    ```

    Example:

    ```bash
    ./BlogAggregator addfeed "TechCrunch" "[https://techcrunch.com/feed/](https://techcrunch.com/feed/)"
    ```

### Follow Feed

-   Allows a user to follow a specific feed.

    ```bash
    ./BlogAggregator follow "<feed_url>"
    ```

    Example:

    ```bash
    ./BlogAggregator follow "[https://techcrunch.com/feed/](https://techcrunch.com/feed/)"
    ```

### Unfollow Feed

-   Allows a user to unfollow a specific feed.

    ```bash
    ./BlogAggregator unfollow "<feed_url>"
    ```

    Example:

    ```bash
    ./BlogAggregator unfollow "[https://techcrunch.com/feed/](https://techcrunch.com/feed/)"
    ```

### List Following Feeds

-   Lists all feeds that the current user is following.

    ```bash
    ./BlogAggregator following
    ```

### Browse Posts

-   Displays the latest posts for the current user.

    ```bash
    ./BlogAggregator browse <number_of_posts>
    ```

    Example:

    ```bash
    ./BlogAggregator browse 3
    ```

### Aggregate Feeds

-   Aggregates posts from followed feeds at specified intervals.

    ```bash
    ./BlogAggregator agg <interval>
    ```

    Example:

    ```bash
    ./BlogAggregator agg 10s
    ```

    Note: The interval must be specified with a unit (e.g., `10s`, `1m`).

### List Feeds

-   Lists all available feeds.

    ```bash
    ./BlogAggregator feeds
    ```

### List Users

-   Lists all registered users.

    ```bash
    ./BlogAggregator users
    ```

### Login

-   Sets the current user for subsequent commands.

    ```bash
    ./BlogAggregator login <username>
    ```

    Example:

    ```bash
    ./BlogAggregator login abc
    ```

## Error Handling

The application provides error messages for common issues, such as:

-   Duplicate feed URLs.
-   Invalid command usage.
-   Database connection errors.
-   Invalid duration format for aggregation.
-   User not logged in for commands that require login.

## Code Structure

-   `main.go`: Entry point of the application, handles command-line arguments and dispatches commands.
-   `internal/config/config.go`: Handles configuration reading from the `.env` file.
-   `internal/database/`: Contains database interaction logic.

## Dependencies

-   `github.com/lib/pq`: PostgreSQL driver.
-   `github.com/joho/godotenv`: For reading environment variables from `.env` files.

## Example Workflow

1.  Register a user:

    ```bash
    ./BlogAggregator register user1
    ```

2.  Login as the user:

    ```bash
    ./BlogAggregator login user1
    ```

3.  Add feeds:

    ```bash
    ./BlogAggregator addfeed "TechCrunch" "[https://techcrunch.com/feed/](https://techcrunch.com/feed/)"
    ./BlogAggregator addfeed "Hacker News" "[https://news.ycombinator.com/rss](https://news.ycombinator.com/rss)"
    ```

4.  Follow feeds:

    ```bash
    ./BlogAggregator follow "[https://techcrunch.com/feed/](https://techcrunch.com/feed/)"
    ./BlogAggregator follow "[https://news.ycombinator.com/rss](https://news.ycombinator.com/rss)"
    ```

5.  Aggregate feeds:

    ```bash
    ./BlogAggregator agg 10s
    ```

6.  Browse posts:

    ```bash
    ./BlogAggregator browse 5
    ```

This README provides a comprehensive guide for using the `BlogAggregator` CLI tool.
