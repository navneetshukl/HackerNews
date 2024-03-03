# HackerNews Concurrency

HackerNews Concurrency is a Go package designed to fetch and display the top stories from Hacker News using concurrency. It leverages goroutines to concurrently fetch multiple stories, enhancing performance and efficiency.

## Overview

Hacker News is a popular news aggregation website where users submit and upvote stories. The HackerNews Concurrency package provides a convenient way to fetch and display the top stories from Hacker News in Go applications. It offers features such as concurrent fetching, caching, and data parsing to streamline the process.
## Features

### Concurrent Fetching

The package utilizes goroutines to fetch multiple stories concurrently, significantly reducing the time taken to retrieve data from Hacker News. By leveraging concurrency, it maximizes performance and responsiveness.

### Caching

To optimize performance and reduce the number of API calls to Hacker News, the package implements a caching mechanism. Fetched stories are cached for a specified duration, ensuring that subsequent requests for the same data can be served from the cache without hitting the Hacker News API again.

### Data Parsing

The package includes functions to parse and modify fetched data for better presentation. It extracts relevant information such as story titles, URLs, and hosts, making it easier to display the data in a user-friendly format.
