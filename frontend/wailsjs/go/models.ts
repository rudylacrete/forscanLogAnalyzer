export namespace models {
	
	export class ForscanLogs {
	    fields: string[];
	    values: number[][];
	
	    static createFrom(source: any = {}) {
	        return new ForscanLogs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fields = source["fields"];
	        this.values = source["values"];
	    }
	}

}

