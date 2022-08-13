import http from "@/http-common"

class APiService {
	getAll(): Promise<any> {
		return http.get("/api/");
	}

	create(data: any): Promise<any> {
		return http.post("/api/new", data);
	}

	update(id: any, data: any): Promise<any> {
		return http.put(`/api/update/${id}`);
	}

	delete(id: any): Promise<any> {
		return http.delete(`/api/delete/${id}`);
	}
}

export default new APiService();