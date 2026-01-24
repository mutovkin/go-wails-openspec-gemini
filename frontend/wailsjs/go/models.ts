export namespace domain {
	
	export class Ticket {
	    Numbers: number[];
	
	    static createFrom(source: any = {}) {
	        return new Ticket(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Numbers = source["Numbers"];
	    }
	}

}

