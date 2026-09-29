export namespace store {
	
	export class Product {
	    ID: number;
	    Name: string;
	    Code: string;
	    Barcode: string;
	    Description: string;
	    Price: number;
	
	    static createFrom(source: any = {}) {
	        return new Product(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Code = source["Code"];
	        this.Barcode = source["Barcode"];
	        this.Description = source["Description"];
	        this.Price = source["Price"];
	    }
	}
	export class UpdateProductParams {
	    Name: string;
	    Code: string;
	    Barcode: string;
	    Description: string;
	    Price: number;
	    ID: number;
	
	    static createFrom(source: any = {}) {
	        return new UpdateProductParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Code = source["Code"];
	        this.Barcode = source["Barcode"];
	        this.Description = source["Description"];
	        this.Price = source["Price"];
	        this.ID = source["ID"];
	    }
	}

}

