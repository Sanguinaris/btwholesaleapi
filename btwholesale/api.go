package btwholesale

import (
	"encoding/json"
)

func BTBroadBandChecker(postCode, town, refNumber, street, buildingNumber, buildingName, districtId string) BTWholeBoi {
	var boi BTWholeBoi
	data := decryptBtWhole(getData("request2", postCode, town, refNumber, street, buildingNumber, buildingName, districtId))
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

func LookupAddress(postCode, town, street, buildingNumber, buildingName string) AddressLookup {
	var addy AddressLookup
	data := decryptBtWhole(getData("request1", postCode, town, "", street, buildingNumber, buildingName, ""))
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
