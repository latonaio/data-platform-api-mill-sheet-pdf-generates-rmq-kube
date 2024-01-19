package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-mill-sheet-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
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

	//	var mechanicalProperties []MechanicalProperty
	//	for _, prop := range headerData.MechanicalProperties {
	//		mechanicalProperties = append(mechanicalProperties, MechanicalProperty{
	//			LabelEn: prop.LabelEn,
	//			Label:   prop.Label,
	//			Unit:    prop.Unit,
	//			Min:     prop.Min,
	//			Max:     prop.Max,
	//			HeatNo:  prop.HeatNo,
	//		})
	//	}

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

	//	var chemicalCompositions []ChemicalComposition
	//	for _, comp := range headerData.ChemicalCompositions {
	//		chemicalCompositions = append(chemicalCompositions, ChemicalComposition{
	//			Label:  comp.Label,
	//			Max:    comp.Max,
	//			Min:    comp.Min,
	//			HeatNo: comp.HeatNo,
	//		})
	//	}

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

	//var specifications []Specification
	//for _, spec := range headerData.Specifications {
	//	specifications = append(specifications, Specification{
	//		LabelEn: spec.LabelEn,
	//		Label:   spec.Label,
	//		Value1:  spec.Value1,
	//		Value2:  spec.Value2,
	//	})
	//}

	pm = &Header{
		OrderID:   headerData.OrderID,
		OrderItem: headerData.OrderItem,
		//		↑OrderNo:							headerData.OrderNo,
		BuyerName: headerData.BuyerName,
		//		↑Customer:							headerData.Customer,
		SellerName:              headerData.SellerName,
		Product:                 headerData.Product,
		SizeOrDimensionText:     headerData.SizeOrDimensionText,
		OrderQuantityInBaseUnit: headerData.OrderQuantityInBaseUnit,
		//		↑Quantity:							headerData.Quantity,
		ProductSpecification: headerData.ProductSpecification,
		//		↑StandardOfRawMaterial:     		headerData.StandardOfRawMaterial,
		MarkingOfMaterial:   headerData.MarkingOfMaterial,
		ProductionOrder:     headerData.ProductionOrder,
		ProductionOrderItem: headerData.ProductionOrderItem,
		//		↑MfgNo:								headerData.MfgNo,
		InspectionLot: headerData.InspectionLot,
		//		↑CertificateNo:						headerData.CertificateNo,
		InspectionPlantBusinessPartnerName: headerData.InspectionPlantBusinessPartnerName,
		//		↑RawMaterialMaker:          		headerData.RawMaterialMaker,
		InspectionLotDate: inputHeaderInspectionLot.InspectionLotDate,
		//		↑Date:								headerData.Date,
		InspectionSpecification: inputHeaderInspectionLot.InspectionSpecification,
		//		↑SpecificationForInspection: 		headerData.SpecificationForInspection,
		HeatNumber:      headerData.HeatNumber,
		DrawingNo:       headerData.DrawingNo,
		PurchaseOrderNo: headerData.PurchaseOrderNo,
		SpecDetails:     specDetails,
		//		MechanicalProperties:       		mechanicalProperties,
		ComponentCompositions: componentCompositions,
		//		ChemicalCompositions:       		chemicalCompositions,
		Inspections: inspections,
		//		Specifications:             		specifications,
		Remarks:                  headerData.Remarks,
		ChiefOfInspectionSection: headerData.ChiefOfInspectionSection,
	}

	return pm
}
