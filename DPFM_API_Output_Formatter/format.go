package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-mill-sheet-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
	"fmt"
)

func SetToPdfData(
	headerData *dpfm_api_input_reader.Header,
	inputHeaderInspectionLot *dpfm_api_input_reader.HeaderInspectionLot,
	inputSpecDetails []dpfm_api_input_reader.SpecDetails,
	inputComponentCompositions []dpfm_api_input_reader.ComponentCompositions,
	inputInspections []dpfm_api_input_reader.Inspections,
) *Header {
	pm := &Header{}

	var specDetails []SpecDetails
	for _, dataSDs := range inputSpecDetails {
		if headerData.OrderID == dataSDs.OrderID {
			specDetails = append(specDetails, SpecDetails{
				OrderID:          dataSDs.OrderID,
				OrderItem:        dataSDs.OrderItem,
				InspectionLot:    dataSDs.InspectionLot,
				SpecType:         dataSDs.SpecType,
				UpperLimitValue:  dataSDs.UpperLimitValue,
				LowerLimitValue:  dataSDs.LowerLimitValue,
				StandardValue:    dataSDs.StandardValue,
				SpecTypeUnit:     dataSDs.SpecTypeUnit,
				HeatNumber:       dataSDs.HeatNumber,
				SpecTypeTextInJA: dataSDs.SpecTypeTextInJA,
			})
		}
	}

	var componentCompositions []ComponentCompositions
	for _, dataCCs := range inputComponentCompositions {
		if headerData.OrderID == dataCCs.OrderID {
			componentCompositions = append(componentCompositions, ComponentCompositions{
				OrderID:                                    dataCCs.OrderID,
				OrderItem:                                  dataCCs.OrderItem,
				InspectionLot:                              dataCCs.InspectionLot,
				ComponentCompositionType:                   dataCCs.ComponentCompositionType,
				ComponentCompositionUpperLimitInPercent:    dataCCs.ComponentCompositionUpperLimitInPercent,
				ComponentCompositionLowerLimitInPercent:    dataCCs.ComponentCompositionLowerLimitInPercent,
				ComponentCompositionStandardValueInPercent: dataCCs.ComponentCompositionStandardValueInPercent,
				HeatNumber:                                 dataCCs.HeatNumber,
			})
		}
	}

	var inspections []Inspections
	for _, dataIPTs := range inputInspections {
		if headerData.OrderID == dataIPTs.OrderID {
			inspections = append(inspections, Inspections{
				OrderID:                                  dataIPTs.OrderID,
				OrderItem:                                dataIPTs.OrderItem,
				InspectionLot:                            dataIPTs.InspectionLot,
				Inspection:                               dataIPTs.Inspection,
				InspectionType:                           dataIPTs.InspectionType,
				InspectionTypeCertificateValueInText:     dataIPTs.InspectionTypeCertificateValueInText,
				InspectionTypeCertificateValueInQuantity: dataIPTs.InspectionTypeCertificateValueInQuantity,
				InspectionTypeValueUnit:                  dataIPTs.InspectionTypeValueUnit,
				InspectionTypeTextInJA:                   dataIPTs.InspectionTypeTextInJA,
			})
		}
	}

	pm = &Header{
		OrderID:                 fmt.Sprintf("%d", headerData.OrderID),
		OrderItem:               headerData.OrderItem,
		BuyerName:               headerData.BuyerName,
		SellerName:              headerData.SellerName,
		Product:                 headerData.Product,
		SizeOrDimensionText:     headerData.SizeOrDimensionText,
		OrderQuantityInBaseUnit: headerData.OrderQuantityInBaseUnit,
		ProductSpecification:    headerData.ProductSpecification,
		MarkingOfMaterial:       headerData.MarkingOfMaterial,
		ProductionOrder:         headerData.ProductionOrder,
		ProductionOrderItem:     headerData.ProductionOrderItem,
		//		↑MfgNo:								headerData.MfgNo,
		InspectionLot: fmt.Sprintf("%d", headerData.InspectionLot),
		//		↑CertificateNo:						headerData.CertificateNo,
		InspectionPlantBusinessPartnerName: headerData.InspectionPlantBusinessPartnerName,
		InspectionLotDate:                  inputHeaderInspectionLot.InspectionLotDate,
		InspectionSpecification:            inputHeaderInspectionLot.InspectionSpecification,
		HeatNumber:                         headerData.HeatNumber,
		DrawingNo:                          headerData.DrawingNo,
		PurchaseOrderNo:                    headerData.PurchaseOrderNo,
		SpecDetails:                        specDetails,
		ComponentCompositions:              componentCompositions,
		Inspections:                        inspections,
		Remarks:                            headerData.Remarks,
		ChiefOfInspectionSection:           headerData.ChiefOfInspectionSection,
	}

	return pm
}
