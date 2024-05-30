import { Microservice, NatsBroker } from "nats-micro";
import kommoMicroservice from "./src/services/kommo.service"

async function main() {
const broker = await new NatsBroker({servers: "nats://nats:4222"}).connect();
await Microservice.create(broker, kommoMicroservice.config);

}

main()