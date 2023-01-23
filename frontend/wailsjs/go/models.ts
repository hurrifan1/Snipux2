export namespace main {
	
	export class Snippet {
	    description: string;
	    name: string;
	    lastModifiedGMT: number;
	    language: string;
	    code: string;
	
	    static createFrom(source: any = {}) {
	        return new Snippet(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.name = source["name"];
	        this.lastModifiedGMT = source["lastModifiedGMT"];
	        this.language = source["language"];
	        this.code = source["code"];
	    }
	}
	export class AppConfig {
	    NumOfLinesToPreview: number;
	    PrimaryColor: string;
	    ColorMode: number;
	
	    static createFrom(source: any = {}) {
	        return new AppConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.NumOfLinesToPreview = source["NumOfLinesToPreview"];
	        this.PrimaryColor = source["PrimaryColor"];
	        this.ColorMode = source["ColorMode"];
	    }
	}

}

