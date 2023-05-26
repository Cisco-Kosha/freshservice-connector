# Kosha Freshservice Connector

Freshservice is a cloud-based customer support platform most commonly used in IT service management functions, enabling IT teams to create service desk tasks, manage projects, and track and report IT assets.

The Kosha Freshservice connector enables you to perform REST API operations from the Freshservice API in your Kosha workflow or custom application. Using the Kosha Freshservice connector, you can directly access the Freshservice platform to:

* List and deactivate Freshservice agents
* Get and delete assets
* List, create, and delete Freshservice tickets

## Useful Actions

You can use the Kosha Freshservice connector to manage Freshservice agents, assets, tickets, and problems. 

Refer to the Kosha Freshservice connector [API specification](openapi.json) for details.

### Agents

Agents are managers and technicians you've added to your service desk. Using the agents API, you can list and deactivate Freshservice agents.

### Assets

IT assets are typically computers, peripherals, other devices, and the software that runs on them. Using the assets API, you can list and delete assets. 

### Tickets

In Freshservice, tickets are objects that enable you to track incidents and service requests. Use the tickets API to list, create, and delete Freshservice tickets.

### Groups

In Freshservice, groups are how you add and organize agents within a workspace. Use the groups API to get details about groups. 

### Problems

In Freshservice, problems specify the root cause of an incident. Use the problems API to list and delete problems in your Freshservice incident data. 

## Authentication

To authenticate when provisioning the Kosha Freshservice connector, you need your:

* Freshservice API key
* Freshservice domain name

## Kosha Connector Open Source Development

All connectors Kosha shares on the marketplace are open source. We believe in fostering collaboration and open development. Everyone is welcome to contribute their ideas, improvements, and feedback for any Kosha connector. We encourage community engagement and appreciate any contributions that align with our goals of an open and collaborative API management platform.

Refer to the contribution guidelines for details.
