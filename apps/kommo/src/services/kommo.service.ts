import type {MicroserviceConfig, Request, Response} from "nats-micro"

class KommoMicroservice {
    public get config(): MicroserviceConfig {
      return {
        name: 'kommo',
        description: 'Simple echo microservice',
        version: '0.0.1',
        methods: {
          say: { handler: this.say },
          'config-change-event': { handler: this.onConfigChange },
        },
      };
    }
  
    private say(req: Request<string>, res: Response<string>): void {
      res.send(req.data);
    }
  
    private onConfigChange(req: Request<unknown>): void {
      console.log(req.data);
    }
  }
  
  export default new KommoMicroservice();