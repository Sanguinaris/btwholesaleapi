package btwholesale

import (
	"encoding/json"
	"os"
)

func BTBroadBandChecker() BTWholeBoi {
	var boi BTWholeBoi
	data := decryptBtWhole(getData("request2", os.Getenv("postCode"), os.Getenv("town"), os.Getenv("refNumber"), os.Getenv("street"), os.Getenv("buildingNumber"), os.Getenv("buildingName"), os.Getenv("districtId")))
	json.Unmarshal(data, &boi)
	return boi
}

func LookupUPRN(uprn string) BTWholeBoi {
	var boi BTWholeBoi
	data := decryptBtWhole(getDataCustomLookup("uprn", uprn))
	json.Unmarshal(data, &boi)
	return boi
}

func LookupAlId(alid string) BTWholeBoi {
	var boi BTWholeBoi
	data := decryptBtWhole(getDataCustomLookup("alid", alid))
	json.Unmarshal(data, &boi)
	return boi
}

func LookupTelephone(telno string) BTWholeBoi {
	var boi BTWholeBoi
	data := decryptBtWhole(getDataCustomLookup("telno", telno))
	json.Unmarshal(data, &boi)
	return boi
}

func LookupAddress(postCode string) AddressLookup {
	var addy AddressLookup
	data := decryptBtWhole(getData("request1", postCode, "", "", "", "", "", ""))
	json.Unmarshal(data, &addy)
	return addy
}

func getData(requestType, postCode, town, refNumber, street, buildingNumber, buildingName, districtId string) string {
	client, req := request_setup()
	req.Header.Add("requestType", requestType)
	req.Header.Add("postCode", postCode)
	req.Header.Add("town", town)
	req.Header.Add("refNumber", refNumber)
	req.Header.Add("street", street)
	req.Header.Add("buildingNumber", buildingNumber)
	req.Header.Add("buildingName", buildingName)
	req.Header.Add("districtId", districtId)
	return request_breakdown(client, req)
}

func getDataCustomLookup(requestType, data string) string {
	client, req := request_setup()
	req.Header.Add("requestType", requestType)
	req.Header.Add(requestType, data)
	return request_breakdown(client, req)
}
