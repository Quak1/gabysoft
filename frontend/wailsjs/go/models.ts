export namespace store {
	
	export class AddTransactionItemParams {
	    Quantity: number;
	    ProductID: number;
	    TransactionID: number;
	
	    static createFrom(source: any = {}) {
	        return new AddTransactionItemParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Quantity = source["Quantity"];
	        this.ProductID = source["ProductID"];
	        this.TransactionID = source["TransactionID"];
	    }
	}
	export class Address {
	    ID: number;
	    City: string;
	    State: string;
	    Country: string;
	    PostalCode: string;
	    Street: string;
	    StreeNumber: string;
	    PhoneNumber: string;
	    TaxID: string;
	    Email: string;
	
	    static createFrom(source: any = {}) {
	        return new Address(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.City = source["City"];
	        this.State = source["State"];
	        this.Country = source["Country"];
	        this.PostalCode = source["PostalCode"];
	        this.Street = source["Street"];
	        this.StreeNumber = source["StreeNumber"];
	        this.PhoneNumber = source["PhoneNumber"];
	        this.TaxID = source["TaxID"];
	        this.Email = source["Email"];
	    }
	}
	export class Client {
	    ID: number;
	    Name: string;
	
	    static createFrom(source: any = {}) {
	        return new Client(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	    }
	}
	export class CreateAddressParams {
	    City: string;
	    State: string;
	    Country: string;
	    PostalCode: string;
	    Street: string;
	    StreeNumber: string;
	    PhoneNumber: string;
	    TaxID: string;
	    Email: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateAddressParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.City = source["City"];
	        this.State = source["State"];
	        this.Country = source["Country"];
	        this.PostalCode = source["PostalCode"];
	        this.Street = source["Street"];
	        this.StreeNumber = source["StreeNumber"];
	        this.PhoneNumber = source["PhoneNumber"];
	        this.TaxID = source["TaxID"];
	        this.Email = source["Email"];
	    }
	}
	export class CreateProductParams {
	    Name: string;
	    Code: string;
	    Barcode: string;
	    Description: string;
	    Price: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateProductParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Code = source["Code"];
	        this.Barcode = source["Barcode"];
	        this.Description = source["Description"];
	        this.Price = source["Price"];
	    }
	}
	export class CreateTransactionParams {
	    // Go type: time
	    Date: any;
	    ClientID: number;
	    AddressID: number;
	    Folio: number;
	    FolioType: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateTransactionParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Date = this.convertValues(source["Date"], null);
	        this.ClientID = source["ClientID"];
	        this.AddressID = source["AddressID"];
	        this.Folio = source["Folio"];
	        this.FolioType = source["FolioType"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Folio {
	    Type: string;
	    Count: number;
	
	    static createFrom(source: any = {}) {
	        return new Folio(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Type = source["Type"];
	        this.Count = source["Count"];
	    }
	}
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
	export class Transaction {
	    ID: number;
	    // Go type: time
	    Date: any;
	    Folio: number;
	    ClientID: number;
	    AddressID: number;
	    FolioType: string;
	
	    static createFrom(source: any = {}) {
	        return new Transaction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Date = this.convertValues(source["Date"], null);
	        this.Folio = source["Folio"];
	        this.ClientID = source["ClientID"];
	        this.AddressID = source["AddressID"];
	        this.FolioType = source["FolioType"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace tasks {
	
	export class TransactionItem {
	    ID: number;
	    Name: string;
	    Code: string;
	    Barcode: string;
	    Description: string;
	    Price: number;
	    Quantity: number;
	
	    static createFrom(source: any = {}) {
	        return new TransactionItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Code = source["Code"];
	        this.Barcode = source["Barcode"];
	        this.Description = source["Description"];
	        this.Price = source["Price"];
	        this.Quantity = source["Quantity"];
	    }
	}
	export class TransactionInfo {
	    ID: number;
	    // Go type: time
	    Date: any;
	    Folio: number;
	    FolioType: string;
	    Client: store.Client;
	    Address: store.Address;
	
	    static createFrom(source: any = {}) {
	        return new TransactionInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Date = this.convertValues(source["Date"], null);
	        this.Folio = source["Folio"];
	        this.FolioType = source["FolioType"];
	        this.Client = this.convertValues(source["Client"], store.Client);
	        this.Address = this.convertValues(source["Address"], store.Address);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class FullTransaction {
	    Info: TransactionInfo;
	    Items: TransactionItem[];
	
	    static createFrom(source: any = {}) {
	        return new FullTransaction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Info = this.convertValues(source["Info"], TransactionInfo);
	        this.Items = this.convertValues(source["Items"], TransactionItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

