# To-Do

The Doozer whiteboard To-Do application. A way of building visible checklists for whatever tasks a Doozer
has on his or her plate at any time. 

## Features

+ User entity
    + Full Name
    + User Name
    + Email
    + Other contact info?
+ Checklist entity
    + Checklist fields
        + Task [255 characters]
        + Priority
            + High
            + Medium
            + Low
        + Status
            + Incomplete
            + Complete
            + Archived
        + Create Date
        + Due Date
        + Complete Date
        + Team?
        + Project?
    + Should a checklist be a User only entity or can a Team and/or Project have a checklist as well?
+ Team entity
    + Users can be added to multiple teams
+ Project entity
    + Users can be added to multiple projects
    + Teams can be added to multiple projects

## Inspirations

+ https://github.com/blue-jay/blueprint
+ https://github.com/josephspurrier/gowebapp

## Development TODO ##

+ Implement a generic approach for handling routes, so that we don't have to define each route.
    + Use gorilla, regex, or URL pattern
+ Research template approaches and implement a standard approach for templating.
    + e.g. Template Inheritance
+ Cleanup code organization and move packages/source around to cut down on the number of imports, etc.
