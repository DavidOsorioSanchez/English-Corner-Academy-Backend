<h1>Documentation</h1>

<p>This is the documentation for the English Corner Academy API.</p>

<h2>API Endpoints</h2>

<p>The API endpoints are documented using the <a href="https://swagger.io/specification/">OpenAPI Specification</a>.</p>

<h3>Authentication</h3>

<p>Authentication is done using the <code>Authorization</code> header with the following format:</p>

<pre><code>Authorization: Bearer &lt;token&gt;
</code></pre>

<p>where <code>&lt;token&gt;</code> is a valid JWT token.</p>

<h3>Events</h3>

<p>The <code>Events</code> endpoint is used to manage events.</p>

<h4>Create an Event</h4>

<p>To create an event, you need to send a <code>POST</code> request to the <code>/events</code> endpoint with the following JSON payload:</p>

<pre><code>{
  "name": "My Event",
  "description": "This is my event",
  "location": "My Location",
  "date": "2023-01-01"
}
</code></pre>

<p>The response will be a JSON object with the following structure:</p>

<pre><code>{
  "event": {
    "id": 1,
    "name": "My Event",
    "description": "This is my event",
    "location": "My Location",
    "date": "2023-01-01",
    "ownerId": 1
  }
}
</code></pre>

<h4>Update an Event</h4>

<p>To update an event, you need to send a <code>PUT</code> request to the <code>/events/:id</code> endpoint with the following JSON payload:</p>

<pre><code>{
  "name": "My Updated Event",
  "description": "This is my updated event",
  "location": "My Updated Location",
  "date": "2023-01-02"
}
</code></pre>

<p>The response will be a JSON object with the following structure:</p>

<pre><code>{
  "event": {
    "id": 1,
    "name": "My Updated Event",
    "description": "This is my updated event",
    "location": "My Updated Location",
    "date": "2023-01-02",
    "ownerId": 1
  }
}
</code></pre>

<h4>Delete an Event</h4>

<p>To delete an event, you need to send a <code>DELETE</code> request to the <code>/events/:id</code> endpoint.</p>

<p>The response will be a JSON object with the following structure:</p>

<pre><code>{
  "message": "Event deleted successfully"
}
</code></pre>

<h4>Get All Events</h4>

<p>To get all events, you need to send a <code>GET</code> request to the <code>/events</code> endpoint.</p>

<p>The response will be a JSON object with the following structure:</p>

<pre><code>{
  "events": [
    {
      "id": 1,
      "name": "My Event",
      "description": "This is my event",
      "location": "My Location",
      "date": "2023-01-01",
      "ownerId": 1
    },
    {
      "id": 2,
      "name": "My Second Event",
      "description": "This is my second event",
      "location": "My Second Location",
      "date": "2023-01-02",
      "ownerId": 1
    }
  ]
}
</code></pre>

<h4>Get an Event by ID</h4>

<p>To get an event by its ID, you need to send a <code>GET</code> request to the <code>/events/:id</code> endpoint.</p>

<p>The response will be a JSON object with the following structure:</p>

<pre><code>{
  "event": {
    "id": 1,
    "name": "My Event",
    "description": "This is my event",
    "location": "My Location",
    "date": "2023-01-01",
    "ownerId": 1
  }
}
</code></pre>

<h4>Get Events by Attendee</h4>

<p>To get events by an attendee, you need to send a <code>GET</code> request to the <code>/events/:id/attendees</code> endpoint.</p>

<p>The response will be a JSON object with the following structure:</p>

<pre><code>{
  "events": [
    {
      "id": 1,
      "name": "My Event",
      "description": "This is my event",
      "location": "My Location",
      "date": "2023-01-01",
      "ownerId": 1
    },
    {
      "id": 2,
      "name": "My Second Event",
      "description": "This is my second event",
      "location": "My Second Location",
      "date": "2023-01-02",
      "ownerId": 1
    }
  ]
}
</code></pre>

<h3>Attendees</h3>

<p>The <code>Attendees</code> endpoint is used to manage attendees.</p>

<h4>Add an Attendee to an Event</h4>

<p>To add an attendee to an event, you need to send a <code>POST</code> request to the <code>/events/:id/attendees/:userId</code> endpoint with the following JSON payload:</p>

<pre><code>{
  "userId": 1
}
</code></pre>

<p>The response will be a JSON object with the following structure:</p>

<pre><code>{
  "message": "Attendee added successfully"
}
</code></pre>

<h4>Get Attendees for an Event</h4>

<p>To get the attendees for an event, you need to send a <code>GET</code> request to the <code>/events/:id/attendees</code> endpoint.</p>

<p>The response will be a JSON object with the following structure:</p>

<pre><code>[
  {
    "id": 1,
    "userId": 1,
    "eventId": 1
  }
]
</code></pre>

<h4>Delete an Attendee from an Event</h4>

<p>To delete an attendee from an event, you need to send a <code>DELETE</code> request to the <code>/events/:id/attendees/:userId</code> endpoint.</p>

<p>The response will be a JSON object with the following structure:</p>

<pre><code>{
  "message": "Attendee deleted successfully"
}
</code></pre>

<h4>Get Attendees by Attendee</h4>

<p>To get attendees by an attendee, you need to send a <code>GET</code> request to the <code>/attendees/:id/events</code> endpoint.</p>

<p>The response will be a JSON object with the following structure:</p>

<pre><code>[
  {
    "id": 1,
    "userId": 1,
    "eventId": 1
  }
]
</code></pre>
